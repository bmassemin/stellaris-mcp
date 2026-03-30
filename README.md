# Stellaris MCP

MCP (Model Context Protocol) server for analyzing Stellaris save games from Claude Desktop.

Parses `.sav` files (ZIP containing a `gamestate` in Clausewitz format) and exposes 10 analysis tools.

## Available Tools

| Tool | Description |
|---|---|
| `get_empire_overview` | Overview: ethics, civics, government, power ratings, resources |
| `get_research_status` | Current research, available options, completed techs |
| `get_economy_breakdown` | Income/expenses by category, net balance |
| `get_planets` | Colonized planets (summary or detail with districts/buildings/modifiers) |
| `get_planets available=true` | Uncolonized habitable planets |
| `get_fleet_power` | Fleets (summary or per-ship detail with weapons and HP) |
| `get_leaders` | Leaders with class, traits, assignment (council/fleet/governor) |
| `get_neighbors` | Known empires: opinion, trust, relative power |
| `get_traditions_ascension` | Adopted traditions and ascension perks |
| `get_notifications` | Wars, federations, players, game status |

All tools accept `country_id` (default 0 = player).
Drill-down tools (`get_planets`, `get_fleet_power`, `get_leaders`) accept an ID for detailed view.

## Installation

### Prerequisites

- Go 1.22+
- Claude Desktop

### Build

```bash
# Linux / macOS
go build -o bin/stellaris-mcp ./cmd/stellaris-mcp

# Windows (cross-compile from Linux/macOS)
GOOS=windows GOARCH=amd64 go build -o bin/stellaris-mcp.exe ./cmd/stellaris-mcp
```

### Claude Desktop Configuration

Add to `claude_desktop_config.json`:

The save games directory is passed as a required CLI argument. The server recursively walks all subdirectories and picks the most recently modified `.sav` file.

**Windows** (`%APPDATA%\Claude\claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "stellaris": {
      "command": "C:\\path\\to\\stellaris-mcp.exe",
      "args": ["C:\\Users\\YourName\\Documents\\Paradox Interactive\\Stellaris\\save games"]
    }
  }
}
```

**macOS** (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "stellaris": {
      "command": "/path/to/stellaris-mcp",
      "args": ["/home/user/.local/share/Paradox Interactive/Stellaris/save games"]
    }
  }
}
```

## Sample Outputs

See [sample_outputs/examples.md](sample_outputs/examples.md) for real output examples of every tool.

## Tests

Copy a `gamestate` file (extracted from a `.sav`) into `internal/clausewitz/testdata/`:

```bash
cp ~/gamestate internal/clausewitz/testdata/gamestate
go test ./...
```

## Architecture

```
cmd/stellaris-mcp/        Entrypoint (stdio MCP server)
internal/
├── clausewitz/           Clausewitz parser (lexer -> parser -> decoder)
│                         Unmarshal(data, &struct) like encoding/json
├── gamestate/            Go structs mapping the Stellaris gamestate
│                         Load from file, ZIP, or save directory
└── tools/                MCP tools (one file per tool)
```
