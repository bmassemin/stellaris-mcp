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

func registerEspionage(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("get_espionage",
			mcp.WithDescription("Get spy networks: infiltration level, spymaster, active operations for each target empire"),
			mcp.WithNumber("country_id", mcp.Description("Country ID (default 0 = player)"), mcp.DefaultNumber(0)),
		),
		handleEspionage,
	)
}

func handleEspionage(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	gs, err := loadLatestSave()
	if err != nil {
		return toolError(err), nil
	}
	countryID, c, err := getCountry(gs, req)
	if err != nil {
		return toolError(err), nil
	}

	type networkEntry struct {
		id  int
		net gamestate.SpyNetwork
	}

	// Our networks (where we spy on others)
	var ourNets []networkEntry
	// Networks targeting us
	var theirNets []networkEntry

	for id, net := range gs.SpyNetworks {
		if net.Owner == countryID {
			ourNets = append(ourNets, networkEntry{id, net})
		}
		if net.Target == countryID {
			theirNets = append(theirNets, networkEntry{id, net})
		}
	}

	sort.Slice(ourNets, func(i, j int) bool { return ourNets[i].net.Power > ourNets[j].net.Power })
	sort.Slice(theirNets, func(i, j int) bool { return theirNets[i].net.Power > theirNets[j].net.Power })

	var b strings.Builder
	fmt.Fprintf(&b, "=== Espionage (%s) ===\n", c.Adjective.Display())

	// Our spy networks
	fmt.Fprintf(&b, "\nOur Spy Networks (%d):\n", len(ourNets))
	activeCount := 0
	for _, e := range ourNets {
		if e.net.Power > 0 || e.net.Leader != 0 {
			activeCount++
		}
	}
	if activeCount == 0 {
		fmt.Fprintf(&b, "  No active spy networks. Assign envoys as spymasters to begin infiltration.\n")
	}
	for _, e := range ourNets {
		target := countryName(gs, e.net.Target)
		if e.net.Power == 0 && e.net.Leader == 0 {
			continue // Skip inactive networks
		}
		fmt.Fprintf(&b, "\n  Target: %s (Country %d)\n", target, e.net.Target)
		fmt.Fprintf(&b, "    Infiltration: %.1f\n", e.net.Power)

		// Resolve spymaster
		if e.net.Leader != 0 {
			if ldr, ok := gs.Leaders[e.net.Leader]; ok {
				fmt.Fprintf(&b, "    Spymaster: %s (Level %d)\n", ldr.Name.FullNames.Display(), ldr.Level)
			}
		}

		if len(e.net.ActiveOperations) > 0 {
			fmt.Fprintf(&b, "    Active Operations: %d\n", len(e.net.ActiveOperations))
		}
		if e.net.Formed != "" {
			fmt.Fprintf(&b, "    Established: %s\n", e.net.Formed)
		}

		// What operations are available at this infiltration level
		writeAvailableOps(&b, e.net.Power)
	}

	// Show inactive networks as a compact list
	inactiveCount := len(ourNets) - activeCount
	if inactiveCount > 0 {
		fmt.Fprintf(&b, "\n  Inactive networks (%d): ", inactiveCount)
		var names []string
		for _, e := range ourNets {
			if e.net.Power == 0 && e.net.Leader == 0 {
				names = append(names, countryName(gs, e.net.Target))
			}
		}
		fmt.Fprintf(&b, "%s\n", strings.Join(names, ", "))
	}

	// Networks targeting us
	fmt.Fprintf(&b, "\nHostile Spy Networks Targeting Us (%d):\n", len(theirNets))
	hostileActive := 0
	for _, e := range theirNets {
		if e.net.Power == 0 && e.net.Leader == 0 {
			continue
		}
		hostileActive++
		owner := countryName(gs, e.net.Owner)
		fmt.Fprintf(&b, "  %s: infiltration %.1f", owner, e.net.Power)
		if len(e.net.ActiveOperations) > 0 {
			fmt.Fprintf(&b, " (%d active ops!)", len(e.net.ActiveOperations))
		}
		fmt.Fprintf(&b, "\n")
	}
	if hostileActive == 0 {
		fmt.Fprintf(&b, "  No known hostile spy networks.\n")
	}

	return mcp.NewToolResultText(b.String()), nil
}

func writeAvailableOps(b *strings.Builder, infiltration float64) {
	type op struct {
		level float64
		name  string
	}
	ops := []op{
		{10, "Gather Information"},
		{20, "Sabotage"},
		{30, "Prepare Sleeper Cells / Acquire Asset"},
		{40, "Steal Technology"},
		{45, "Sabotage Starbase"},
		{60, "Arm Privateers"},
		{70, "Spark Diplomatic Incident"},
		{80, "Crisis Beacon"},
	}
	var available, locked []string
	for _, o := range ops {
		if infiltration >= o.level {
			available = append(available, fmt.Sprintf("%s (%.0f)", o.name, o.level))
		} else {
			locked = append(locked, fmt.Sprintf("%s (%.0f)", o.name, o.level))
		}
	}
	if len(available) > 0 {
		fmt.Fprintf(b, "    Available Ops: %s\n", strings.Join(available, ", "))
	}
	if len(locked) > 0 {
		fmt.Fprintf(b, "    Next unlock: %s\n", locked[0])
	}
}
