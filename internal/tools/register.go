package tools

import "github.com/mark3labs/mcp-go/server"

func Register(s *server.MCPServer, dir string) {
	saveDir = dir
	registerPing(s)
	registerEmpireOverview(s)
	registerResearchStatus(s)
	registerNeighbors(s)
	registerFleetPower(s)
	registerEconomyBreakdown(s)
	registerPlanets(s)
	registerTraditionsAscension(s)
	registerNotifications(s)
	registerLeaders(s)
}
