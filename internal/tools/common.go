package tools

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bmassemin/stellaris-mcp/internal/gamestate"
	"github.com/mark3labs/mcp-go/mcp"
)

var (
	saveDir string
	loc     *gamestate.Localizer
)

func loadLatestSave() (*gamestate.GameState, error) {
	return gamestate.LoadFromDir(saveDir)
}

// l resolves a game key to its localized display name.
func l(key string) string {
	return loc.Resolve(key)
}

// ll resolves a slice of game keys.
func ll(keys []string) []string {
	return loc.ResolveAll(keys)
}

func getCountry(gs *gamestate.GameState, req mcp.CallToolRequest) (int, *gamestate.Country, error) {
	id := int(req.GetFloat("country_id", 0))
	c, ok := gs.Country[id]
	if !ok {
		return 0, nil, fmt.Errorf("country %d not found", id)
	}
	return id, &c, nil
}

func countryName(gs *gamestate.GameState, id int) string {
	if c, ok := gs.Country[id]; ok {
		return c.Adjective.Display()
	}
	return fmt.Sprintf("Country %d", id)
}

func formatResources(r map[string]float64) string {
	if len(r) == 0 {
		return "(none)"
	}
	keys := make([]string, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s: %.1f", k, r[k]))
	}
	return strings.Join(parts, ", ")
}

func toolError(err error) *mcp.CallToolResult {
	return mcp.NewToolResultError(err.Error())
}
