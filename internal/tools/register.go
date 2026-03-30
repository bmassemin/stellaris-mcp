package tools

import (
	"github.com/bmassemin/stellaris-mcp/internal/gamestate"
	"github.com/mark3labs/mcp-go/server"
)

func Register(s *server.MCPServer, dir string, localizer *gamestate.Localizer) {
	saveDir = dir
	loc = localizer
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
