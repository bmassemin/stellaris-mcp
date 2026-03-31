package main

import (
	"fmt"
	"os"

	"github.com/bmassemin/stellaris-mcp/internal/gamestate"
	"github.com/bmassemin/stellaris-mcp/internal/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: stellaris-mcp <save-games-directory> <localization-directory>")
		os.Exit(1)
	}
	saveDir := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: stellaris-mcp <save-games-directory> <localization-directory>")
		os.Exit(1)
	}
	locDir := os.Args[2]
	if _, err := os.Stat(locDir); err != nil {
		fmt.Fprintf(os.Stderr, "localization directory not found: %s\n", locDir)
		os.Exit(1)
	}
	loc := gamestate.NewLocalizer(locDir)
	if loc.Loaded() == 0 {
		fmt.Fprintf(os.Stderr, "no localization entries found in: %s\n", locDir)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "stellaris-mcp: loaded %d localization entries\n", loc.Loaded())

	s := server.NewMCPServer(
		"stellaris-mcp",
		"0.1.0",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)

	tools.Register(s, saveDir, loc)

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
