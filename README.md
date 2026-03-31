# Stellaris MCP

MCP (Model Context Protocol) server for analyzing Stellaris save games from Claude Desktop or Claude Code.

Parses `.sav` files (ZIP containing a `gamestate` in Clausewitz format) and exposes 11 analysis tools. Includes a `/stellaris` skill with compiled strategy guides for patch 4.3 (Cetus).

## Available Tools

| Tool | Description |
|---|---|
| `get_empire_overview` | Ethics, civics, government, power ratings, naval/starbase capacity, resources |
| `get_research_status` | Current research + progress, available options, completed techs |
| `get_economy_breakdown` | Income/expenses by category, net balance per resource |
| `get_planets` | Colonized planets with districts, buildings, jobs, modifiers, production |
| `get_planets available=true` | Surveyed uncolonized habitable planets with features |
| `get_fleet_power` | Fleets with power, ship-level detail (design, HP, weapons) |
| `get_leaders` | Leaders with class, traits, council position, assignment |
| `get_neighbors` | Known empires: opinion, trust, relative power comparison |
| `get_traditions_ascension` | Adopted traditions and ascension perks |
| `get_espionage` | Spy networks: infiltration, spymasters, available operations |
| `get_notifications` | Wars, federations, players, game status |

All tools accept `country_id` (default 0 = player).

Drill-down tools (`get_planets`, `get_fleet_power`, `get_leaders`) accept an ID for detailed view.

## `/stellaris` Skill (Claude Code)

A Claude Code skill at `.claude/skills/stellaris/` provides strategic advice by combining live save data with compiled strategy guides:

- **Economy** -- districts, specializations, trade, designations, mega-economy
- **Military** -- fleet composition, ship design, war timing, crisis counter-builds
- **Research** -- tech priorities, traditions, ascension perks, beelining
- **Expansion** -- colonization criteria, chokepoints, wide vs tall
- **Diplomacy** -- federations, vassals, leaders, council, espionage

Usage: `/stellaris how should I optimize my economy?`

## Installation

### Prerequisites

- Go 1.22+
- Claude Desktop or Claude Code

### Build

```bash
# Linux / macOS
go build -o bin/stellaris-mcp ./cmd/stellaris-mcp

# Windows (cross-compile from Linux/macOS)
GOOS=windows GOARCH=amd64 go build -o bin/stellaris-mcp.exe ./cmd/stellaris-mcp
```

Or download a pre-built binary from [Releases](https://github.com/bmassemin/stellaris-mcp/releases).

### Usage

```
stellaris-mcp <save-games-directory> <localization-directory>
```

- **save-games-directory** (required): root folder containing your Stellaris saves. The server recursively walks all subdirectories and picks the most recently modified `.sav` file.
- **localization-directory** (required): path to the `english` localization folder from your Stellaris installation. Provides real in-game names (e.g. "Energy Credits", "Forge Capital"). The server exits with an error if the path doesn't exist or contains no localization files.

### Claude Desktop

Add to `claude_desktop_config.json`:

**Windows** (`%APPDATA%\Claude\claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "stellaris": {
      "command": "C:\\path\\to\\stellaris-mcp.exe",
      "args": [
        "C:\\Users\\YourName\\Documents\\Paradox Interactive\\Stellaris\\save games",
        "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Stellaris\\localisation\\english"
      ]
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
      "args": [
        "/home/user/.local/share/Paradox Interactive/Stellaris/save games",
        "/home/user/.steam/steam/steamapps/common/Stellaris/localisation/english"
      ]
    }
  }
}
```

### Claude Code

```bash
claude mcp add --transport stdio stellaris \
  -- /path/to/stellaris-mcp \
  "/path/to/Stellaris/save games" \
  "/path/to/Stellaris/localisation/english"
```

The `/stellaris` skill is automatically available when working in this repository.

## Sample Outputs

See [sample_outputs/examples.md](sample_outputs/examples.md) for real output examples of every tool, generated from a mid-game save with English localization.

## Tests

Copy a `gamestate` file (extracted from a `.sav`) into `internal/clausewitz/testdata/`:

```bash
cp ~/gamestate internal/clausewitz/testdata/gamestate
```

To run localization tests, create a `.env` file at the project root (see `.env.example`):

```bash
cp .env.example .env
# Edit .env with your Stellaris localisation/english path
```

Then run:

```bash
go test ./...
```

Tests that require the fixture or localization files skip gracefully when they are absent.

## Architecture

```
cmd/stellaris-mcp/           Entry point (stdio MCP server)
internal/
├── clausewitz/              Clausewitz format parser (lexer -> parser -> decoder)
│                            Unmarshal(data, &struct) like encoding/json
├── gamestate/               Go structs mapping the Stellaris gamestate
│                            Load from raw bytes, ZIP, or save directory
│                            Localization service (YAML loc files -> l(key))
├── tools/                   MCP tools (one file per tool)
│                            common.go: shared helpers (loadLatestSave, getCountry, l, ll)
└── testutil/                .env file loader for tests
.claude/skills/stellaris/    /stellaris skill + strategy guides
```
