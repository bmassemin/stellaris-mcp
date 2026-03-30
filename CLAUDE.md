# Stellaris MCP

## Build & Test
- `go build ./...` — build all packages
- `go build -o bin/stellaris-mcp ./cmd/stellaris-mcp` — build binary
- `GOOS=windows GOARCH=amd64 go build -o bin/stellaris-mcp.exe ./cmd/stellaris-mcp` — cross-compile for Windows
- `go test ./... -timeout 120s` — run all tests (fixture parsing takes ~0.15s)
- Test fixture (`internal/clausewitz/testdata/gamestate`) is gitignored (15MB); copy from a real save to run tests

## Architecture
- `internal/clausewitz/` — Clausewitz format parser. `Unmarshal(data, &struct)` like `encoding/json`. Tag: `clausewitz:"field_name"`.
- `internal/gamestate/` — Go structs mapping Stellaris gamestate. `Load(data)`, `LoadFromZip(path)`, `LoadFromDir(dir)`.
- `internal/tools/` — MCP tools. Pattern: `registerXxx(s)` + `handleXxx(ctx, req)`. One file per tool.
- `cmd/stellaris-mcp/` — Entry point. CLI: `stellaris-mcp <save-dir> [loc-dir]` — loc-dir points to Steam's `localisation/english/` folder.
- `internal/gamestate/localization.go` — Loads Stellaris YAML loc files. `l(key)`/`ll(keys)` in tools resolve keys to display names, fallback to `PrettyKey`.

## Key Gotchas
- Clausewitz maps contain `none` entries (e.g. `1=none` in districts) — decoder skips them in `decodeMap`
- Planet `Owner` field defaults to `0` (Go zero value) which is also country 0 (player). Always check `isColonized()` (pops > 0 or designation != "") when filtering owned planets
- Stellaris names use templates (`PLANET_NAME_FORMAT`, `STAR_NAME_X_OF_Y`, `%ADJECTIVE%`, `$var$`) — resolved recursively in `LocalizedName.Display()`
- `SPEC_*` name keys are game localization references not present in save files; they cannot be resolved further
- Save files are `.sav` ZIPs containing a `gamestate` file; `LoadFromDir` walks subdirs recursively for most recent `.sav`
- For large responses (planets list), use summary/detail pattern with ID parameter to stay under Claude Desktop 1MB limit
- Localization YAML format: ` key:version "value"`. Some values are `$concept_X$` references — `Resolve()` follows one level of indirection
- Don't name variables `l` in tool handlers — it shadows the `l()` localization helper
- After changing any tool output, regenerate `sample_outputs/examples.md` using the sample_test.go generator pattern (write it, run it, delete it)

## Adding a New Tool
1. Create `internal/tools/myfeature.go` with `registerMyFeature(s)` + `handleMyFeature(ctx, req)`
2. Call `registerMyFeature(s)` in `internal/tools/register.go`
3. Use `loadLatestSave()` and `getCountry(gs, req)` from `common.go`
4. Return `mcp.NewToolResultText(text)` or `toolError(err)`
5. Use `l(key)` to localize game keys (ethics, techs, buildings, resources, etc.)
