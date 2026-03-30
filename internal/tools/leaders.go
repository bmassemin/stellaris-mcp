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

func registerLeaders(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("get_leaders",
			mcp.WithDescription("List leaders with class, assignment, level, traits. Pass leader_id for full detail."),
			mcp.WithNumber("country_id", mcp.Description("Country ID (default 0 = player)"), mcp.DefaultNumber(0)),
			mcp.WithNumber("leader_id", mcp.Description("Leader ID for detailed view (omit for summary)"), mcp.DefaultNumber(-1)),
		),
		handleLeaders,
	)
}

func handleLeaders(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	gs, err := loadLatestSave()
	if err != nil {
		return toolError(err), nil
	}

	leaderID := int(req.GetFloat("leader_id", -1))
	if leaderID >= 0 {
		return leaderDetail(gs, leaderID)
	}

	countryID, c, err := getCountry(gs, req)
	if err != nil {
		return toolError(err), nil
	}
	return leaderSummary(gs, countryID, c)
}

func leaderSummary(gs *gamestate.GameState, countryID int, c *gamestate.Country) (*mcp.CallToolResult, error) {
	type entry struct {
		id int
		l  gamestate.Leader
	}
	owned := make(map[int]bool, len(c.OwnedLeaders))
	for _, id := range c.OwnedLeaders {
		owned[id] = true
	}
	var leaders []entry
	for id, ldr := range gs.Leaders {
		if owned[id] {
			leaders = append(leaders, entry{id, ldr})
		}
	}
	sort.Slice(leaders, func(i, j int) bool { return leaders[i].l.Class < leaders[j].l.Class })

	var b strings.Builder
	fmt.Fprintf(&b, "=== Leaders (%s) — %d leaders ===\n", c.Adjective.Display(), len(leaders))
	fmt.Fprintf(&b, "Use get_leaders with leader_id for full detail.\n\n")

	fmt.Fprintf(&b, "%-12s %-25s %-12s %3s %3s %-20s %s\n",
		"ID", "Name", "Class", "Lvl", "Age", "Assignment", "Traits")
	fmt.Fprintf(&b, "%s\n", strings.Repeat("-", 110))

	for _, e := range leaders {
		assignment := resolveAssignment(gs, e.l)
		traits := strings.Join(ll(e.l.Traits), ", ")
		if traits == "" {
			traits = "-"
		}
		fmt.Fprintf(&b, "%-12d %-25s %-12s %3d %3d %-20s %s\n",
			e.id,
			truncate(e.l.Name.FullNames.Display(), 25),
			e.l.Class,
			e.l.Level,
			e.l.Age,
			truncate(assignment, 20),
			traits,
		)
	}

	if len(leaders) == 0 {
		fmt.Fprintf(&b, "No leaders.\n")
	}

	return mcp.NewToolResultText(b.String()), nil
}

func leaderDetail(gs *gamestate.GameState, leaderID int) (*mcp.CallToolResult, error) {
	ldr, ok := gs.Leaders[leaderID]
	if !ok {
		return toolError(fmt.Errorf("leader %d not found", leaderID)), nil
	}

	var b strings.Builder
	fmt.Fprintf(&b, "=== Leader %d: %s ===\n", leaderID, ldr.Name.FullNames.Display())
	fmt.Fprintf(&b, "Class: %s\n", ldr.Class)
	fmt.Fprintf(&b, "Tier: %s\n", ldr.Tier)
	fmt.Fprintf(&b, "Level: %d (bonus: %d)\n", ldr.Level, ldr.BonusSkillLevel)
	fmt.Fprintf(&b, "Experience: %.1f\n", ldr.Experience)
	fmt.Fprintf(&b, "Age: %d\n", ldr.Age)
	fmt.Fprintf(&b, "Gender: %s\n", ldr.Gender)
	fmt.Fprintf(&b, "Ethic: %s\n", l(ldr.Ethic))
	fmt.Fprintf(&b, "Job: %s\n", l(ldr.Job))
	fmt.Fprintf(&b, "Recruited: %s\n", ldr.RecruitmentDate)

	fmt.Fprintf(&b, "\nAssignment: %s\n", resolveAssignment(gs, ldr))

	if ldr.Location.Type != "" {
		fmt.Fprintf(&b, "Location: %s (id=%d)\n", ldr.Location.Type, ldr.Location.ID)
	}
	if ldr.CouncilLocation.Type != "" {
		fmt.Fprintf(&b, "Council: %s (id=%d, position=%d)\n",
			ldr.CouncilLocation.Type, ldr.CouncilLocation.ID, ldr.CouncilLocation.Position)
	}

	fmt.Fprintf(&b, "\nTraits:\n")
	if len(ldr.Traits) == 0 {
		fmt.Fprintf(&b, "  (none)\n")
	}
	for _, t := range ldr.Traits {
		fmt.Fprintf(&b, "  - %s\n", l(t))
	}

	return mcp.NewToolResultText(b.String()), nil
}

func resolveAssignment(gs *gamestate.GameState, ldr gamestate.Leader) string {
	// Council position
	if ldr.CouncilLocation.Type == "council_position" {
		return fmt.Sprintf("Council #%d", ldr.CouncilLocation.ID)
	}
	// Fleet/ship assignment
	if ldr.Location.Type == "ship" {
		if fleet, ok := gs.Fleet[ldr.Location.ID]; ok {
			return fmt.Sprintf("Fleet: %s", fleet.Name.Display())
		}
		return fmt.Sprintf("Ship %d", ldr.Location.ID)
	}
	// Governor
	if ldr.Location.Type == "planet" {
		if p, ok := gs.Planets.Planet[ldr.Location.ID]; ok {
			return fmt.Sprintf("Governor: %s", p.Name.Display())
		}
	}
	if ldr.Job != "" {
		return l(ldr.Job)
	}
	return "Unassigned"
}
