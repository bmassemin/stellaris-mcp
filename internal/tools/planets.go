package tools

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/bmassemin/stellaris-mcp/internal/gamestate"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerPlanets(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("get_planets",
			mcp.WithDescription("List owned planets (summary). Use available=true to list uncolonized habitable planets surveyed by this empire. Pass planet_id for full details."),
			mcp.WithNumber("country_id", mcp.Description("Country ID (default 0 = player)"), mcp.DefaultNumber(0)),
			mcp.WithNumber("planet_id", mcp.Description("Planet ID for detailed view (omit for summary list)"), mcp.DefaultNumber(-1)),
			mcp.WithBoolean("available", mcp.Description("Show uncolonized habitable planets instead of owned"), mcp.DefaultBool(false)),
		),
		handlePlanets,
	)
}

func handlePlanets(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	gs, err := loadLatestSave()
	if err != nil {
		return toolError(err), nil
	}
	countryID, c, err := getCountry(gs, req)
	if err != nil {
		return toolError(err), nil
	}

	planetID := int(req.GetFloat("planet_id", -1))
	if planetID >= 0 {
		return planetDetail(gs, planetID)
	}
	if req.GetBool("available", false) {
		return planetAvailable(gs, countryID, c)
	}
	return planetSummary(gs, countryID, c)
}

func planetSummary(gs *gamestate.GameState, countryID int, c *gamestate.Country) (*mcp.CallToolResult, error) {
	type entry struct {
		id int
		p  gamestate.Planet
	}
	var owned []entry
	for id, p := range gs.Planets.Planet {
		if p.Owner == countryID && isColonized(p) {
			owned = append(owned, entry{id, p})
		}
	}
	sort.Slice(owned, func(i, j int) bool { return owned[i].id < owned[j].id })

	var b strings.Builder
	fmt.Fprintf(&b, "=== Planets (%s) — %d planets ===\n", c.Adjective.Display(), len(owned))
	fmt.Fprintf(&b, "Use get_planets with planet_id for details.\n\n")
	fmt.Fprintf(&b, "%-6s %-25s %-16s %4s %5s %5s %-12s %s\n", "ID", "Name", "Class", "Size", "Pops", "Stab", "Design.", "Districts")
	fmt.Fprintf(&b, "%s\n", strings.Repeat("-", 110))

	for _, e := range owned {
		fmt.Fprintf(&b, "%-6d %-25s %-16s %4d %5d %5.0f %-12s %s\n",
			e.id,
			truncate(e.p.Name.Display(), 25),
			l(e.p.PlanetClass),
			e.p.PlanetSize,
			e.p.NumPops,
			e.p.Stability,
			l(e.p.FinalDesignation),
			districtSummary(gs, e.p),
		)
	}

	if len(owned) == 0 {
		fmt.Fprintf(&b, "No owned planets.\n")
	}

	return mcp.NewToolResultText(b.String()), nil
}

func planetDetail(gs *gamestate.GameState, planetID int) (*mcp.CallToolResult, error) {
	p, ok := gs.Planets.Planet[planetID]
	if !ok {
		return toolError(fmt.Errorf("planet %d not found", planetID)), nil
	}

	var b strings.Builder
	fmt.Fprintf(&b, "=== Planet %d: %s ===\n", planetID, p.Name.Display())
	fmt.Fprintf(&b, "Class: %s, Size: %d\n", l(p.PlanetClass), p.PlanetSize)
	fmt.Fprintf(&b, "Designation: %s (Ascension Tier %d)\n", l(p.FinalDesignation), p.AscensionTier)
	fmt.Fprintf(&b, "Owner: %d, Controller: %d\n", p.Owner, p.Controller)
	fmt.Fprintf(&b, "\nPopulation: %d pops\n", p.NumPops)
	fmt.Fprintf(&b, "Stability: %.1f\n", p.Stability)
	fmt.Fprintf(&b, "Crime: %.1f\n", p.Crime)
	fmt.Fprintf(&b, "Amenities: %.0f (used: %.0f, free: %.0f)\n", p.Amenities, p.AmenitiesUsage, p.FreeAmenities)
	fmt.Fprintf(&b, "Housing: %.0f (used: %.0f, free: %.0f)\n", p.TotalHousing, p.HousingUsage, p.FreeHousing)

	// Jobs breakdown
	jobCounts := jobBreakdown(gs, planetID)
	if len(jobCounts) > 0 {
		fmt.Fprintf(&b, "\nJobs:\n")
		for _, jc := range jobCounts {
			fmt.Fprintf(&b, "  %s: %d\n", l(jc.jobType), jc.workforce)
		}
	}

	// District slot summary
	counts := countDistrictTypes(gs, p)
	totalUsed := 0
	for _, n := range counts {
		totalUsed += n
	}
	fmt.Fprintf(&b, "\nDistrict Slots: %d used (planet size %d)\n", totalUsed, p.PlanetSize)
	for dtype, n := range counts {
		fmt.Fprintf(&b, "  %s: %d\n", l(dtype), n)
	}

	if len(p.Districts) > 0 {
		fmt.Fprintf(&b, "\nDistricts Detail (%d):\n", len(p.Districts))
		for _, did := range p.Districts {
			d, ok := gs.Districts[did]
			if !ok {
				continue
			}
			lvl := ""
			if d.Level > 1 {
				lvl = fmt.Sprintf(" (lvl %d)", d.Level)
			}
			// Resolve zone specializations
			var specs []string
			for _, zid := range d.Zones {
				if z, ok := gs.Zones[zid]; ok && z.Type != "zone_default" {
					specs = append(specs, l(z.Type))
				}
			}
			if len(specs) > 0 {
				fmt.Fprintf(&b, "  %s%s — slots: %s\n", l(d.Type), lvl, strings.Join(specs, ", "))
			} else {
				fmt.Fprintf(&b, "  %s%s\n", l(d.Type), lvl)
			}
			// Buildings inside this district's zones
			for _, zid := range d.Zones {
				z, ok := gs.Zones[zid]
				if !ok {
					continue
				}
				for _, bid := range z.Buildings {
					if bld, ok := gs.Buildings[bid]; ok {
						fmt.Fprintf(&b, "    [%s] %s\n", l(z.Type), l(bld.Type))
					}
				}
			}
		}
	}

	// Planetary deposits (features, blockers)
	if len(p.Deposits) > 0 {
		var features, blockers []string
		for _, did := range p.Deposits {
			dep, ok := gs.Deposit[did]
			if !ok || dep.Type == "" {
				continue
			}
			if isBlocker(dep.Type) {
				label := l(dep.Type)
				if dep.SwapType != "" {
					label += " -> " + l(dep.SwapType)
				}
				blockers = append(blockers, label)
			} else {
				features = append(features, l(dep.Type))
			}
		}
		if len(features) > 0 {
			fmt.Fprintf(&b, "\nPlanetary Features (%d):\n", len(features))
			for _, f := range features {
				fmt.Fprintf(&b, "  - %s\n", f)
			}
		}
		if len(blockers) > 0 {
			fmt.Fprintf(&b, "\nBlockers (%d):\n", len(blockers))
			for _, bl := range blockers {
				fmt.Fprintf(&b, "  - %s\n", bl)
			}
		}
	}

	// Timed modifiers
	if len(p.TimedModifier.Items) > 0 {
		fmt.Fprintf(&b, "\nModifiers:\n")
		for _, m := range p.TimedModifier.Items {
			if m.Days < 0 {
				fmt.Fprintf(&b, "  - %s (permanent)\n", l(m.Modifier))
			} else {
				fmt.Fprintf(&b, "  - %s (%d days remaining)\n", l(m.Modifier), m.Days)
			}
		}
	}

	if len(p.Produces) > 0 {
		fmt.Fprintf(&b, "\nMonthly Production:\n  %s\n", formatResources(p.Produces))
	}
	if len(p.Upkeep) > 0 {
		fmt.Fprintf(&b, "Monthly Upkeep:\n  %s\n", formatResources(p.Upkeep))
	}
	if len(p.Profits) > 0 {
		fmt.Fprintf(&b, "Monthly Profit:\n  %s\n", formatResources(p.Profits))
	}

	return mcp.NewToolResultText(b.String()), nil
}

func countDistrictTypes(gs *gamestate.GameState, p gamestate.Planet) map[string]int {
	counts := make(map[string]int)
	for _, did := range p.Districts {
		if d, ok := gs.Districts[did]; ok {
			level := d.Level
			if level <= 0 {
				level = 1
			}
			counts[d.Type] += level
		}
	}
	return counts
}

func districtSummary(gs *gamestate.GameState, p gamestate.Planet) string {
	counts := countDistrictTypes(gs, p)
	if len(counts) == 0 {
		return fmt.Sprintf("0/%d", p.PlanetSize)
	}
	totalUsed := 0
	for _, n := range counts {
		totalUsed += n
	}
	parts := make([]string, 0, len(counts))
	for dtype, n := range counts {
		parts = append(parts, fmt.Sprintf("%s:%d", l(dtype), n))
	}
	sort.Strings(parts)
	return fmt.Sprintf("%d [%s]", totalUsed, strings.Join(parts, " "))
}

type jobCount struct {
	jobType   string
	workforce int
}

func jobBreakdown(gs *gamestate.GameState, planetID int) []jobCount {
	counts := make(map[string]int)
	for _, job := range gs.PopJobs {
		if job.Planet == planetID && job.Workforce > 0 {
			counts[job.Type] += job.Workforce
		}
	}
	result := make([]jobCount, 0, len(counts))
	for t, w := range counts {
		result = append(result, jobCount{t, w})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].workforce > result[j].workforce })
	return result
}

func depositSummary(gs *gamestate.GameState, p gamestate.Planet) string {
	if len(p.Deposits) == 0 {
		return "-"
	}
	var names []string
	for _, did := range p.Deposits {
		dep, ok := gs.Deposit[did]
		if !ok || dep.Type == "" {
			continue
		}
		names = append(names, l(dep.Type))
	}
	if len(names) == 0 {
		return "-"
	}
	return strings.Join(names, ", ")
}

var habitableClasses = map[string]bool{
	"pc_continental": true, "pc_tropical": true, "pc_arid": true,
	"pc_desert": true, "pc_ocean": true, "pc_arctic": true,
	"pc_alpine": true, "pc_savannah": true, "pc_tundra": true,
	"pc_gaia": true, "pc_tomb": true, "pc_relic": true,
}

func isColonized(p gamestate.Planet) bool {
	return p.NumPops > 0 || p.FinalDesignation != ""
}

func planetAvailable(gs *gamestate.GameState, countryID int, c *gamestate.Country) (*mcp.CallToolResult, error) {
	type entry struct {
		id       int
		p        gamestate.Planet
		surveyed bool
	}
	var planets []entry
	for id, p := range gs.Planets.Planet {
		if !habitableClasses[p.PlanetClass] || isColonized(p) {
			continue
		}
		// Skip planets owned by another empire
		if p.Owner != 0 {
			continue
		}
		if p.SurveyedBy != countryID {
			continue
		}
		planets = append(planets, entry{id, p, true})
	}
	// Sort by size descending
	sort.Slice(planets, func(i, j int) bool {
		return planets[i].p.PlanetSize > planets[j].p.PlanetSize
	})

	var b strings.Builder
	fmt.Fprintf(&b, "=== Available Habitable Planets (%s) — %d planets ===\n", c.Adjective.Display(), len(planets))
	fmt.Fprintf(&b, "Use get_planets with planet_id for details.\n\n")

	fmt.Fprintf(&b, "%-6s %-25s %-18s %4s %s\n", "ID", "Name", "Class", "Size", "Features")
	fmt.Fprintf(&b, "%s\n", strings.Repeat("-", 90))

	for _, e := range planets {
		fmt.Fprintf(&b, "%-6d %-25s %-18s %4d %s\n",
			e.id,
			truncate(e.p.Name.Display(), 25),
			l(e.p.PlanetClass),
			e.p.PlanetSize,
			depositSummary(gs, e.p),
		)
	}

	if len(planets) == 0 {
		fmt.Fprintf(&b, "No available habitable planets found.\n")
	}

	return mcp.NewToolResultText(b.String()), nil
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-1] + "…"
}

func isBlocker(depositType string) bool {
	blockers := []string{
		"_blocker",
		"d_failing_infrastructure",
		"d_decrepit_dwellings",
		"d_active_volcano",
		"d_deep_sinkhole",
		"d_dense_jungle",
		"d_noxious_swamp",
		"d_quicksand_basin",
		"d_radioactive_wasteland",
		"d_toxic_kelp",
		"d_machine_empire_ruins",
	}
	for _, b := range blockers {
		if strings.Contains(depositType, b) || depositType == b {
			return true
		}
	}
	return false
}
