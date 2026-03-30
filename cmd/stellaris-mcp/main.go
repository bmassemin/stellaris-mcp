package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bmassemin/stellaris-mcp/internal/gamestate"
	"github.com/bmassemin/stellaris-mcp/internal/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: stellaris-mcp <save-games-directory>")
		os.Exit(1)
	}
	saveDir := os.Args[1]

	// Load localization files from data/english/ next to the binary,
	// or from the working directory.
	exe, _ := os.Executable()
	dataDir := filepath.Join(filepath.Dir(exe), "data", "english")
	if _, err := os.Stat(dataDir); err != nil {
		dataDir = filepath.Join("data", "english")
	}
	loc := gamestate.NewLocalizer(dataDir)
	if n := loc.Loaded(); n > 0 {
		fmt.Fprintf(os.Stderr, "stellaris-mcp: loaded %d localization entries\n", n)
	}

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
