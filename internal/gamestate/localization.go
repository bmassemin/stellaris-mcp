package gamestate

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// Localizer resolves Stellaris internal keys to display names.
type Localizer struct {
	entries map[string]string
}

// NewLocalizer loads all .yml localization files from dir recursively.
// Returns an empty (but usable) Localizer if dir doesn't exist or is empty.
func NewLocalizer(dir string) *Localizer {
	l := &Localizer{entries: make(map[string]string)}
	filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(d.Name(), ".yml") {
			return nil
		}
		l.loadFile(path)
		return nil
	})
	return l
}

// Resolve returns the localized display name for a key.
// Replaces all $reference$ in the value with their resolved entries.
// Falls back to PrettyKey if not found.
func (l *Localizer) Resolve(key string) string {
	v, ok := l.entries[key]
	if !ok {
		return PrettyKey(key)
	}
	if strings.Contains(v, "$") {
		v = l.resolveRefs(v)
	}
	return stripMarkup(v)
}

// resolveRefs replaces all $key$ references in a string.
func (l *Localizer) resolveRefs(s string) string {
	for i := 0; i < 5; i++ { // max depth to avoid infinite loops
		start := strings.Index(s, "$")
		if start < 0 {
			break
		}
		end := strings.Index(s[start+1:], "$")
		if end < 0 {
			break
		}
		end += start + 1
		ref := s[start+1 : end]
		if resolved, ok := l.entries[ref]; ok {
			s = s[:start] + resolved + s[end+1:]
		} else {
			// Skip unresolvable ref — replace $ to avoid infinite loop
			s = s[:start] + ref + s[end+1:]
		}
	}
	return s
}

// ResolveAll resolves a slice of keys.
func (l *Localizer) ResolveAll(keys []string) []string {
	out := make([]string, len(keys))
	for i, k := range keys {
		out[i] = l.Resolve(k)
	}
	return out
}

// Has returns true if the key exists in the localization data.
func (l *Localizer) Has(key string) bool {
	_, ok := l.entries[key]
	return ok
}

// Loaded returns the number of entries loaded.
func (l *Localizer) Loaded() int {
	return len(l.entries)
}

// loadFile parses a Stellaris localization YAML file.
// Format: ` key:version "value"`
func (l *Localizer) loadFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip BOM, empty lines, comments, header (l_english:)
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' || strings.HasPrefix(line, "l_english") {
			continue
		}
		// Format: key:N "value" or key: "value"
		colonIdx := strings.Index(line, ":")
		if colonIdx <= 0 {
			continue
		}
		key := line[:colonIdx]

		// Find the quoted value
		quoteStart := strings.Index(line[colonIdx:], "\"")
		if quoteStart < 0 {
			continue
		}
		quoteStart += colonIdx + 1
		quoteEnd := strings.LastIndex(line, "\"")
		if quoteEnd <= quoteStart {
			continue
		}
		value := line[quoteStart:quoteEnd]

		// Skip descriptions and tooltips
		if strings.HasSuffix(key, "_desc") || strings.HasSuffix(key, "_tt") || strings.HasSuffix(key, "_tooltip") {
			continue
		}

		l.entries[key] = value
	}
}

// stripMarkup cleans Stellaris text markup: £icon£ -> icon name, §X...§! -> remove, \n -> space.
func stripMarkup(s string) string {
	// Replace £icon_name£ with the icon name (e.g. £energy£ -> "Energy")
	for {
		start := strings.Index(s, "\u00a3")
		if start < 0 {
			break
		}
		end := strings.Index(s[start+2:], "\u00a3")
		if end < 0 {
			break
		}
		icon := s[start+2 : start+2+end]
		// Map known icons to readable names; skip pure formatting icons
		replacement := ""
		switch icon {
		case "blocker", "icon", "trigger_no", "trigger_yes":
			// Pure formatting icons, just remove
		default:
			replacement = titleCase(strings.ReplaceAll(icon, "_", " "))
		}
		s = s[:start] + replacement + s[start+2+end+2:]
	}
	// Remove §X (color start) and §! (color end)
	for {
		idx := strings.Index(s, "\u00a7")
		if idx < 0 {
			break
		}
		// §! is 2+1 bytes, §X is 2+1 bytes
		if idx+3 <= len(s) {
			s = s[:idx] + s[idx+3:]
		} else {
			s = s[:idx]
		}
	}
	s = strings.ReplaceAll(s, "\\n", " ")
	return strings.TrimSpace(s)
}
