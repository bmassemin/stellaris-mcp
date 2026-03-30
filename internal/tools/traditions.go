package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerTraditionsAscension(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("get_traditions_ascension",
			mcp.WithDescription("Get adopted traditions and ascension perks"),
			mcp.WithNumber("country_id", mcp.Description("Country ID (default 0 = player)"), mcp.DefaultNumber(0)),
		),
		handleTraditionsAscension,
	)
}

func handleTraditionsAscension(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	gs, err := loadLatestSave()
	if err != nil {
		return toolError(err), nil
	}
	_, c, err := getCountry(gs, req)
	if err != nil {
		return toolError(err), nil
	}

	var b strings.Builder
	fmt.Fprintf(&b, "=== Traditions & Ascension Perks (%s) ===\n", c.Adjective.Display())

	fmt.Fprintf(&b, "\nTraditions:\n")
	if len(c.Traditions) == 0 {
		fmt.Fprintf(&b, "  (none adopted)\n")
	} else {
		for _, t := range c.Traditions {
			fmt.Fprintf(&b, "  - %s\n", l(t))
		}
	}

	fmt.Fprintf(&b, "\nAscension Perks:\n")
	if len(c.AscensionPerks) == 0 {
		fmt.Fprintf(&b, "  (none adopted)\n")
	} else {
		for _, p := range c.AscensionPerks {
			fmt.Fprintf(&b, "  - %s\n", l(p))
		}
	}

	if len(c.Traditions) == 0 && len(c.AscensionPerks) == 0 {
		fmt.Fprintf(&b, "\nNote: Early game — traditions require Unity to adopt.\n")
	}

	return mcp.NewToolResultText(b.String()), nil
}
