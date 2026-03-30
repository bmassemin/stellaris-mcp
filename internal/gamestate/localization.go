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
	return v
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
