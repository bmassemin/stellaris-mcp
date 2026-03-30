package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/bmassemin/stellaris-mcp/internal/gamestate"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerFleetPower(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("get_fleet_power",
			mcp.WithDescription("List fleets with power (summary). Pass fleet_id for ship-level detail with weapons and HP."),
			mcp.WithNumber("country_id", mcp.Description("Country ID (default 0 = player)"), mcp.DefaultNumber(0)),
			mcp.WithNumber("fleet_id", mcp.Description("Fleet ID for detailed view (omit for summary)"), mcp.DefaultNumber(-1)),
		),
		handleFleetPower,
	)
}

func handleFleetPower(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	gs, err := loadLatestSave()
	if err != nil {
		return toolError(err), nil
	}

	fleetID := int(req.GetFloat("fleet_id", -1))
	if fleetID >= 0 {
		return fleetDetail(gs, fleetID)
	}

	_, c, err := getCountry(gs, req)
	if err != nil {
		return toolError(err), nil
	}
	return fleetSummary(gs, c)
}

func fleetSummary(gs *gamestate.GameState, c *gamestate.Country) (*mcp.CallToolResult, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "=== Fleet Power (%s) ===\n", c.Adjective.Display())
	fmt.Fprintf(&b, "Total Military Power: %.1f\n", c.MilitaryPower)
	fmt.Fprintf(&b, "Fleet Size: %d\n", c.FleetSize)
	fmt.Fprintf(&b, "Use get_fleet_power with fleet_id for ship details.\n\n")

	fmt.Fprintf(&b, "%-8s %-30s %-25s %8s %5s\n", "ID", "Name", "Class", "Power", "Ships")
	fmt.Fprintf(&b, "%s\n", strings.Repeat("-", 80))

	milPower := 0.0
	for _, of := range c.FleetsManager.OwnedFleets {
		fleet, ok := gs.Fleet[of.Fleet]
		if !ok {
			continue
		}
		role := l(fleet.ShipClass)
		if fleet.Station {
			role += " (station)"
		}
		if fleet.Civilian {
			role += " (civilian)"
		}
		fmt.Fprintf(&b, "%-8d %-30s %-25s %8.1f %5d\n",
			of.Fleet, truncate(fleet.Name.Display(), 30), role, fleet.MilitaryPower, len(fleet.Ships))
		if !fleet.Station && !fleet.Civilian {
			milPower += fleet.MilitaryPower
		}
	}

	fmt.Fprintf(&b, "\nMilitary Fleet Power (excl. stations/civilian): %.1f\n", milPower)

	// Comparison
	fmt.Fprintf(&b, "\nComparison with Known Empires:\n")
	hasNeighbors := false
	for _, rel := range c.RelationsManager.Relation {
		if !rel.Contact {
			continue
		}
		if neighbor, ok := gs.Country[rel.Country]; ok {
			hasNeighbors = true
			fmt.Fprintf(&b, "  %s: %.1f military power\n", neighbor.Adjective.Display(), neighbor.MilitaryPower)
		}
	}
	if !hasNeighbors {
		fmt.Fprintf(&b, "  No known empires.\n")
	}

	return mcp.NewToolResultText(b.String()), nil
}

func fleetDetail(gs *gamestate.GameState, fleetID int) (*mcp.CallToolResult, error) {
	fleet, ok := gs.Fleet[fleetID]
	if !ok {
		return toolError(fmt.Errorf("fleet %d not found", fleetID)), nil
	}

	var b strings.Builder
	fmt.Fprintf(&b, "=== Fleet %d: %s ===\n", fleetID, fleet.Name.Display())
	fmt.Fprintf(&b, "Class: %s\n", l(fleet.ShipClass))
	fmt.Fprintf(&b, "Military Power: %.1f\n", fleet.MilitaryPower)
	fmt.Fprintf(&b, "Ships: %d\n", len(fleet.Ships))
	if fleet.Station {
		fmt.Fprintf(&b, "Station: yes\n")
	}
	if fleet.Civilian {
		fmt.Fprintf(&b, "Civilian: yes\n")
	}

	for _, sid := range fleet.Ships {
		ship, ok := gs.Ships[sid]
		if !ok {
			continue
		}
		// Resolve design name and ship size
		designName := "unknown"
		shipSize := "unknown"
		if d, ok := gs.ShipDesign[ship.ShipDesignImplementation.Design]; ok {
			designName = d.Name.Display()
			if len(d.GrowthStages) > 0 {
				shipSize = d.GrowthStages[0].ShipSize
			}
		}

		fmt.Fprintf(&b, "\n  --- Ship %d: %s ---\n", sid, ship.Name.Display())
		fmt.Fprintf(&b, "  Design: %s (%s)\n", designName, l(shipSize))
		fmt.Fprintf(&b, "  Hull:   %.0f / %.0f\n", ship.Hitpoints, ship.MaxHitpoints)
		fmt.Fprintf(&b, "  Shield: %.0f / %.0f\n", ship.ShieldHitpoints, ship.MaxShieldHitpoints)
		fmt.Fprintf(&b, "  Armor:  %.0f / %.0f\n", ship.ArmorHitpoints, ship.MaxArmorHitpoints)
		fmt.Fprintf(&b, "  Section: %s\n", l(ship.Section.Design))

		if len(ship.Section.Weapon) > 0 {
			fmt.Fprintf(&b, "  Weapons:\n")
			for _, w := range ship.Section.Weapon {
				fmt.Fprintf(&b, "    - %s\n", l(w.Template))
			}
		}
	}

	return mcp.NewToolResultText(b.String()), nil
}
