package tools

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerEconomyBreakdown(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("get_economy_breakdown",
			mcp.WithDescription("Get detailed income and expenses breakdown by category"),
			mcp.WithNumber("country_id", mcp.Description("Country ID (default 0 = player)"), mcp.DefaultNumber(0)),
		),
		handleEconomyBreakdown,
	)
}

func handleEconomyBreakdown(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	gs, err := loadLatestSave()
	if err != nil {
		return toolError(err), nil
	}
	_, c, err := getCountry(gs, req)
	if err != nil {
		return toolError(err), nil
	}

	month := c.Budget.CurrentMonth
	var b strings.Builder
	fmt.Fprintf(&b, "=== Economy Breakdown (%s) ===\n", c.Adjective.Display())

	// Income
	fmt.Fprintf(&b, "\n--- INCOME ---\n")
	totalIncome := make(map[string]float64)
	writeBudgetSection(&b, month.Income, totalIncome)
	fmt.Fprintf(&b, "  TOTAL: %s\n", formatResources(totalIncome))

	// Expenses
	fmt.Fprintf(&b, "\n--- EXPENSES ---\n")
	totalExpenses := make(map[string]float64)
	writeBudgetSection(&b, month.Expenses, totalExpenses)
	fmt.Fprintf(&b, "  TOTAL: %s\n", formatResources(totalExpenses))

	// Net
	fmt.Fprintf(&b, "\n--- NET BALANCE ---\n")
	net := make(map[string]float64)
	for k, v := range totalIncome {
		net[k] += v
	}
	for k, v := range totalExpenses {
		net[k] -= v
	}
	keys := make([]string, 0, len(net))
	for k := range net {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sign := "+"
		if net[k] < 0 {
			sign = ""
		}
		fmt.Fprintf(&b, "  %s: %s%.1f\n", k, sign, net[k])
	}

	return mcp.NewToolResultText(b.String()), nil
}

func writeBudgetSection(b *strings.Builder, categories map[string]map[string]float64, totals map[string]float64) {
	keys := make([]string, 0, len(categories))
	for k := range categories {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, cat := range keys {
		resources := categories[cat]
		fmt.Fprintf(b, "  %s: %s\n", cat, formatResources(resources))
		for k, v := range resources {
			totals[k] += v
		}
	}
}
