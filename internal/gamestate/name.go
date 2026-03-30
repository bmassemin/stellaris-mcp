package gamestate

import "strings"

// LocalizedName represents a Clausewitz localized string reference
// with optional variable substitutions.
type LocalizedName struct {
	Key       string     `clausewitz:"key"`
	Variables []NameVar  `clausewitz:"variables"`
}

type NameVar struct {
	Key   string        `clausewitz:"key"`
	Value LocalizedName `clausewitz:"value"`
}

// Display resolves the name to a human-readable string.
func (n LocalizedName) Display() string {
	if n.Key == "" {
		return "(unnamed)"
	}

	// Build variable map (recursively resolve values)
	vars := make(map[string]string)
	for _, v := range n.Variables {
		vars[v.Key] = v.Value.Display()
	}

	switch {
	// "Iolam III" — PLANET_NAME_FORMAT with PARENT + NUMERAL
	case n.Key == "PLANET_NAME_FORMAT":
		parent := vars["PARENT"]
		numeral := vars["NUMERAL"]
		if parent != "" && numeral != "" {
			return parent + " " + numeral
		}
		if parent != "" {
			return parent
		}

	// "Bebaki System" — STAR_NAME_X_OF_Y with NAME
	case strings.HasPrefix(n.Key, "STAR_NAME_"):
		if name := vars["NAME"]; name != "" {
			return name
		}

	// "%ADJECTIVE%" or "%ADJ%" — template substitution
	case strings.Contains(n.Key, "%"):
		result := n.Key
		for k, v := range vars {
			result = strings.ReplaceAll(result, "%"+strings.ToUpper(k)+"%", v)
			// Also try lowercase placeholder
			result = strings.ReplaceAll(result, "%"+k+"%", v)
		}
		// If all placeholders resolved, return
		if !strings.Contains(result, "%") {
			return result
		}
		// Partial resolution — try positional
		for _, v := range n.Variables {
			return v.Value.Display()
		}

	// "$affix$$base$" — direct variable concatenation
	case strings.Contains(n.Key, "$"):
		result := n.Key
		for k, v := range vars {
			result = strings.ReplaceAll(result, "$"+k+"$", v)
		}
		if !strings.Contains(result, "$") {
			return result
		}
	}

	// No variables or unrecognized pattern — return key as-is
	if len(vars) == 0 {
		return n.Key
	}

	// Fallback: concatenate variable values
	parts := make([]string, 0, len(vars))
	for _, v := range n.Variables {
		parts = append(parts, v.Value.Display())
	}
	return strings.Join(parts, " ")
}
