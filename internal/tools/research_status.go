package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/bmassemin/stellaris-mcp/internal/gamestate"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerResearchStatus(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("get_research_status",
			mcp.WithDescription("Get current research, available options, and completed technologies"),
			mcp.WithNumber("country_id", mcp.Description("Country ID (default 0 = player)"), mcp.DefaultNumber(0)),
		),
		handleResearchStatus,
	)
}

func handleResearchStatus(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	gs, err := loadLatestSave()
	if err != nil {
		return toolError(err), nil
	}
	_, c, err := getCountry(gs, req)
	if err != nil {
		return toolError(err), nil
	}

	ts := c.TechStatus
	var b strings.Builder
	fmt.Fprintf(&b, "=== Research Status (%s) ===\n", c.Adjective.Display())

	fmt.Fprintf(&b, "\nCurrent Research:\n")
	writeQueue(&b, "Physics", ts.PhysicsQueue)
	writeQueue(&b, "Society", ts.SocietyQueue)
	writeQueue(&b, "Engineering", ts.EngineeringQueue)

	fmt.Fprintf(&b, "\nAvailable Research Options:\n")
	writeAlternatives(&b, "Physics", ts.Alternatives.Physics)
	writeAlternatives(&b, "Society", ts.Alternatives.Society)
	writeAlternatives(&b, "Engineering", ts.Alternatives.Engineering)

	fmt.Fprintf(&b, "\nCompleted Technologies (%d):\n", len(ts.Technology))
	for _, t := range ts.Technology {
		fmt.Fprintf(&b, "  - %s\n", l(t))
	}

	return mcp.NewToolResultText(b.String()), nil
}

func writeQueue(b *strings.Builder, label string, queue []gamestate.ResearchItem) {
	if len(queue) == 0 {
		fmt.Fprintf(b, "  %s: (none)\n", label)
		return
	}
	for _, r := range queue {
		fmt.Fprintf(b, "  %s: %s (progress: %.1f)\n", label, l(r.Technology), r.Progress)
	}
}

func writeAlternatives(b *strings.Builder, label string, alts []string) {
	if len(alts) == 0 {
		fmt.Fprintf(b, "  %s: (none)\n", label)
		return
	}
	fmt.Fprintf(b, "  %s: %s\n", label, strings.Join(ll(alts), ", "))
}
