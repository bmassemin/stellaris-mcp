package gamestate

import (
	"strings"
	"testing"

	"github.com/bmassemin/stellaris-mcp/internal/testutil"
)

func TestLocalizer(t *testing.T) {
	dir := testutil.LoadEnv("STELLARIS_LOC_DIR")
	if dir == "" {
		t.Skip("STELLARIS_LOC_DIR not set in .env, skipping localization test")
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

	// Deposits with markup should be stripped
	t.Run("deposit_markup", func(t *testing.T) {
		got := loc.Resolve("d_energy_5")
		if strings.Contains(got, "\u00a3") || strings.Contains(got, "\u00a7") {
			t.Errorf("markup not stripped: %q", got)
		}
		if got == "" {
			t.Error("resolved to empty string")
		}
		t.Logf("d_energy_5 -> %q", got)
	})

	// Deposit features (no markup)
	t.Run("deposit_feature", func(t *testing.T) {
		got := loc.Resolve("d_hot_springs")
		if got != "Hot Springs" {
			t.Errorf("got %q, want %q", got, "Hot Springs")
		}
	})
}
