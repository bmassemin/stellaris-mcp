package gamestate

import (
	"os"
	"testing"
)

func TestLocalizer(t *testing.T) {
	dir := "../../data/english"
	if _, err := os.Stat(dir); err != nil {
		t.Skip("data/english not found, skipping localization test")
	}

	loc := NewLocalizer(dir)
	if loc.Loaded() == 0 {
		t.Fatal("no entries loaded")
	}
	t.Logf("Loaded %d entries", loc.Loaded())

	tests := []struct {
		key  string
		want string
	}{
		{"ethic_authoritarian", "Authoritarian"},
		{"civic_technocracy", "Technocracy"},
		{"gov_science_directorate", "Science Directorate"},
		{"auth_oligarchic", "Oligarchic"},
		{"tech_administrative_ai", "Administrative AI"},
		{"origin_default", "Prosperous Unification"},
		{"district_city", "City District"},
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			got := loc.Resolve(tc.key)
			if got != tc.want {
				t.Errorf("Resolve(%q) = %q, want %q", tc.key, got, tc.want)
			}
		})
	}

	// Unknown key falls back to PrettyKey
	got := loc.Resolve("some_unknown_key_xyz")
	if got == "some_unknown_key_xyz" {
		t.Error("unknown key should be pretty-printed, not returned raw")
	}
}
