package main

import (
	"fmt"
	"os"

	"github.com/bmassemin/stellaris-mcp/internal/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: stellaris-mcp <save-games-directory>")
		os.Exit(1)
	}
	saveDir := os.Args[1]

	s := server.NewMCPServer(
		"stellaris-mcp",
		"0.1.0",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)

	tools.Register(s, saveDir)

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
