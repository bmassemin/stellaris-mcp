package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerNotifications(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("get_notifications",
			mcp.WithDescription("Get active wars, federations, and game status"),
			mcp.WithNumber("country_id", mcp.Description("Country ID (default 0 = player)"), mcp.DefaultNumber(0)),
		),
		handleNotifications,
	)
}

func handleNotifications(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	gs, err := loadLatestSave()
	if err != nil {
		return toolError(err), nil
	}
	_, c, err := getCountry(gs, req)
	if err != nil {
		return toolError(err), nil
	}

	var b strings.Builder
	fmt.Fprintf(&b, "=== Game Status (%s — %s) ===\n", c.Adjective.Display(), gs.Date)

	// Players
	fmt.Fprintf(&b, "\nPlayers:\n")
	if len(gs.Player) == 0 {
		fmt.Fprintf(&b, "  (none)\n")
	}
	for _, p := range gs.Player {
		fmt.Fprintf(&b, "  - %s (Country %d)\n", p.Name, p.Country)
	}

	// Wars
	fmt.Fprintf(&b, "\nActive Wars:\n")
	if len(gs.War) == 0 {
		fmt.Fprintf(&b, "  No active wars.\n")
	}
	for _, w := range gs.War {
		fmt.Fprintf(&b, "  - %s (started %s)\n", w.Name.Display(), w.StartDate)
	}

	// Federations
	fmt.Fprintf(&b, "\nFederations:\n")
	if len(gs.Federation) == 0 {
		fmt.Fprintf(&b, "  No federations.\n")
	}
	for _, f := range gs.Federation {
		fedType := l(f.FederationProgression.FederationType)
		fmt.Fprintf(&b, "  - %s (%s, leader: %s, %d members, cohesion: %.0f%%)\n",
			f.Name.Display(), fedType, countryName(gs, f.Leader),
			len(f.Members), f.FederationProgression.Cohesion)
	}

	// Known contacts count
	contacts := 0
	for _, rel := range c.RelationsManager.Relation {
		if rel.Contact {
			contacts++
		}
	}
	fmt.Fprintf(&b, "\nDiplomacy:\n")
	fmt.Fprintf(&b, "  Known empires: %d\n", contacts)
	fmt.Fprintf(&b, "  Victory rank: %d\n", c.VictoryRank)

	return mcp.NewToolResultText(b.String()), nil
}
