package gamestate

import (
	"os"
	"testing"
)

func TestLoad_Fixture(t *testing.T) {
	data, err := os.ReadFile("../clausewitz/testdata/gamestate")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	gs, err := Load(data)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if gs.Version == "" {
		t.Error("Version is empty")
	}
	if gs.Date == "" {
		t.Error("Date is empty")
	}
	if len(gs.Player) == 0 {
		t.Fatal("no players")
	}

	// Country 0
	c, ok := gs.Country[0]
	if !ok {
		t.Fatal("Country 0 not found")
	}

	t.Run("ethics", func(t *testing.T) {
		if len(c.Ethos.Ethic) == 0 {
			t.Fatal("no ethics")
		}
	})

	t.Run("government", func(t *testing.T) {
		if c.Government.Type == "" {
			t.Error("government type is empty")
		}
		if c.Government.Authority == "" {
			t.Error("authority is empty")
		}
		if len(c.Government.Civics) == 0 {
			t.Error("no civics")
		}
	})

	t.Run("tech_status", func(t *testing.T) {
		if len(c.TechStatus.Technology) == 0 {
			t.Error("no completed techs")
		}
	})

	t.Run("budget", func(t *testing.T) {
		if len(c.Budget.CurrentMonth.Income) == 0 {
			t.Error("no income categories")
		}
		if len(c.Budget.CurrentMonth.Expenses) == 0 {
			t.Error("no expense categories")
		}
	})

	t.Run("power_ratings", func(t *testing.T) {
		if c.MilitaryPower == 0 {
			t.Error("military power is 0")
		}
	})

	t.Run("fleets_manager", func(t *testing.T) {
		if len(c.FleetsManager.OwnedFleets) == 0 {
			t.Error("no owned fleets")
		}
	})

	t.Run("fleets", func(t *testing.T) {
		if len(gs.Fleet) == 0 {
			t.Fatal("no fleets loaded")
		}
	})

	t.Run("planets", func(t *testing.T) {
		if len(gs.Planets.Planet) == 0 {
			t.Fatal("no planets loaded")
		}
		// Find a colonized planet owned by country 0
		found := false
		for _, p := range gs.Planets.Planet {
			if p.Owner == 0 && p.NumPops > 0 {
				found = true
				break
			}
		}
		if !found {
			t.Error("no colonized planets found for country 0")
		}
	})

	t.Run("districts_buildings", func(t *testing.T) {
		if len(gs.Districts) == 0 {
			t.Fatal("no districts loaded")
		}
		if len(gs.Buildings) == 0 {
			t.Fatal("no buildings loaded")
		}
	})

	t.Run("leaders", func(t *testing.T) {
		if len(gs.Leaders) == 0 {
			t.Fatal("no leaders loaded")
		}
		// Find a leader belonging to country 0
		found := false
		for _, l := range gs.Leaders {
			if l.Country == 0 && l.Class != "" {
				found = true
				break
			}
		}
		if !found {
			t.Error("no leaders found for country 0")
		}
	})
}
