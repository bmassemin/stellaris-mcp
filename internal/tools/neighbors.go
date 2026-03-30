package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerNeighbors(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("get_neighbors",
			mcp.WithDescription("Get known empires: relations, opinion, trust, and relative power"),
			mcp.WithNumber("country_id", mcp.Description("Country ID (default 0 = player)"), mcp.DefaultNumber(0)),
		),
		handleNeighbors,
	)
}

func handleNeighbors(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	gs, err := loadLatestSave()
	if err != nil {
		return toolError(err), nil
	}
	_, c, err := getCountry(gs, req)
	if err != nil {
		return toolError(err), nil
	}

	var b strings.Builder
	fmt.Fprintf(&b, "=== Known Empires (%s) ===\n", c.Adjective.Display())
	fmt.Fprintf(&b, "Your Military Power: %.1f\n\n", c.MilitaryPower)

	contacts := 0
	for _, rel := range c.RelationsManager.Relation {
		if !rel.Contact {
			continue
		}
		contacts++
		neighbor, ok := gs.Country[rel.Country]
		if !ok {
			continue
		}
		ratio := ""
		if c.MilitaryPower > 0 {
			r := neighbor.MilitaryPower / c.MilitaryPower
			switch {
			case r > 2:
				ratio = "Overwhelming"
			case r > 1.5:
				ratio = "Superior"
			case r > 0.75:
				ratio = "Equivalent"
			case r > 0.4:
				ratio = "Inferior"
			default:
				ratio = "Pathetic"
			}
		}
		fmt.Fprintf(&b, "%s (Country %d):\n", neighbor.Adjective.Display(), rel.Country)
		fmt.Fprintf(&b, "  Opinion: %.0f, Trust: %.0f\n", rel.RelationCurrent, rel.Trust)
		fmt.Fprintf(&b, "  Military: %.1f (%s)\n", neighbor.MilitaryPower, ratio)
		fmt.Fprintf(&b, "  Economy: %.1f, Tech: %.1f\n\n", neighbor.EconomyPower, neighbor.TechPower)
	}

	if contacts == 0 {
		fmt.Fprintf(&b, "No known empires yet.\n")
	}

	return mcp.NewToolResultText(b.String()), nil
}
