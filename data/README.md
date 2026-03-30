# Localization Data

Place the Stellaris localization folder here to enable human-readable name resolution.

## Setup

Copy the `localisation` folder from your Stellaris installation:

**Steam (Windows):**
```
C:\Program Files (x86)\Steam\steamapps\common\Stellaris\localisation\
```

**Steam (Linux):**
```
~/.steam/steam/steamapps/common/Stellaris/localisation/
```

**Steam (macOS):**
```
~/Library/Application Support/Steam/steamapps/common/Stellaris/localisation/
```

Into this directory:
```
data/localisation/
```

## What it does

The game save files use internal keys like `SPEC_Bebaki` or `NAME_Aureyon` instead of display names. These keys are resolved using YAML localization files (e.g. `english/name_lists_l_english.yml`).

Without these files, the MCP displays raw keys. With them, it can show proper in-game names.

## Note

These files are copyrighted by Paradox Interactive and must not be committed to this repository.
