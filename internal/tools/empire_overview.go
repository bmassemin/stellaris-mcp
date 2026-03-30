package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerEmpireOverview(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("get_empire_overview",
			mcp.WithDescription("Get empire overview: ethics, civics, government, resources, pops, power ratings"),
			mcp.WithNumber("country_id", mcp.Description("Country ID (default 0 = player)"), mcp.DefaultNumber(0)),
		),
		handleEmpireOverview,
	)
}

func handleEmpireOverview(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	gs, err := loadLatestSave()
	if err != nil {
		return toolError(err), nil
	}
	_, c, err := getCountry(gs, req)
	if err != nil {
		return toolError(err), nil
	}

	// Count owned planets (exclude uncolonized — Owner defaults to 0)
	countryID := int(req.GetFloat("country_id", 0))
	planetCount := 0
	for _, p := range gs.Planets.Planet {
		if p.Owner == countryID && (p.NumPops > 0 || p.FinalDesignation != "") {
			planetCount++
		}
	}

	// Net balance from budget
	netBalance := make(map[string]float64)
	for _, resources := range c.Budget.CurrentMonth.Balance {
		for k, v := range resources {
			netBalance[k] += v
		}
	}

	var b strings.Builder
	fmt.Fprintf(&b, "=== Empire Overview (%s) ===\n", gs.Date)
	fmt.Fprintf(&b, "Game: %s (version %s)\n\n", gs.Name, gs.Version)
	fmt.Fprintf(&b, "Name: %s\n", c.Adjective.Display())
	fmt.Fprintf(&b, "Government: %s (%s)\n", c.Government.Type, c.Government.Authority)
	fmt.Fprintf(&b, "Origin: %s\n", c.Government.Origin)
	fmt.Fprintf(&b, "Ethics: %s\n", strings.Join(c.Ethos.Ethic, ", "))
	fmt.Fprintf(&b, "Civics: %s\n", strings.Join(c.Government.Civics, ", "))
	fmt.Fprintf(&b, "\nPower Ratings:\n")
	fmt.Fprintf(&b, "  Military: %.1f\n", c.MilitaryPower)
	fmt.Fprintf(&b, "  Economy:  %.1f\n", c.EconomyPower)
	fmt.Fprintf(&b, "  Tech:     %.1f\n", c.TechPower)
	fmt.Fprintf(&b, "  Victory Rank: %d\n", c.VictoryRank)
	fmt.Fprintf(&b, "\nEmpire Stats:\n")
	fmt.Fprintf(&b, "  Pops: %d\n", c.NumPops)
	fmt.Fprintf(&b, "  Planets: %d\n", planetCount)
	fmt.Fprintf(&b, "  Fleet Size: %d\n", c.FleetSize)
	fmt.Fprintf(&b, "  Empire Size: %d\n", c.EmpireSize)
	fmt.Fprintf(&b, "\nMonthly Net Balance:\n  %s\n", formatResources(netBalance))

	return mcp.NewToolResultText(b.String()), nil
}
