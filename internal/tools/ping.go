package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerPing(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("ping",
			mcp.WithDescription("Verify that the Stellaris MCP server is running"),
		),
		handlePing,
	)
}

func handlePing(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText("pong"), nil
}
