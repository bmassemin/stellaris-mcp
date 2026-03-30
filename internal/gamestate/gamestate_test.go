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

	// Header
	if gs.Version != "Cetus v4.3.2" {
		t.Errorf("Version = %q", gs.Version)
	}
	if gs.Date != "2200.07.01" {
		t.Errorf("Date = %q", gs.Date)
	}

	// Players
	if len(gs.Player) != 2 {
		t.Fatalf("Player count = %d, want 2", len(gs.Player))
	}
	if gs.Player[0].Name != "Froupix" {
		t.Errorf("Player[0].Name = %q", gs.Player[0].Name)
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
		if len(c.TechStatus.PhysicsQueue) == 0 {
			t.Error("no physics research in queue")
		}
		if len(c.TechStatus.Alternatives.Physics) == 0 {
			t.Error("no physics alternatives")
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
		if c.EconomyPower == 0 {
			t.Error("economy power is 0")
		}
	})

	t.Run("fleets_manager", func(t *testing.T) {
		if len(c.FleetsManager.OwnedFleets) == 0 {
			t.Error("no owned fleets")
		}
	})

	t.Run("relations", func(t *testing.T) {
		// Country 0 may not have relations in early game.
		// Check a country that does (16777219).
		if c2, ok := gs.Country[16777219]; ok {
			if len(c2.RelationsManager.Relation) == 0 {
				t.Error("country 16777219 has no relations")
			}
		}
	})

	// Fleets
	t.Run("fleets", func(t *testing.T) {
		if len(gs.Fleet) == 0 {
			t.Fatal("no fleets loaded")
		}
		// Check fleet 0 (starbase)
		f, ok := gs.Fleet[0]
		if !ok {
			t.Fatal("fleet 0 not found")
		}
		if f.MilitaryPower == 0 {
			t.Error("fleet 0 military power is 0")
		}
	})

	// Planets
	t.Run("planets", func(t *testing.T) {
		if len(gs.Planets.Planet) == 0 {
			t.Fatal("no planets loaded")
		}
		// Check planet 11 (capital)
		p, ok := gs.Planets.Planet[11]
		if !ok {
			t.Fatal("planet 11 not found")
		}
		if p.PlanetClass != "pc_continental" {
			t.Errorf("planet 11 class = %q", p.PlanetClass)
		}
		if p.NumPops == 0 {
			t.Error("planet 11 has 0 pops")
		}
		if p.Stability == 0 {
			t.Error("planet 11 stability is 0")
		}
	})

	// Districts & Buildings resolution
	t.Run("districts_buildings", func(t *testing.T) {
		if len(gs.Districts) == 0 {
			t.Fatal("no districts loaded")
		}
		if len(gs.Buildings) == 0 {
			t.Fatal("no buildings loaded")
		}
		// District 27 = district_generator
		d, ok := gs.Districts[27]
		if !ok {
			t.Fatal("district 27 not found")
		}
		if d.Type != "district_generator" {
			t.Errorf("district 27 type = %q", d.Type)
		}
		// Building 0 = building_capital
		b, ok := gs.Buildings[0]
		if !ok {
			t.Fatal("building 0 not found")
		}
		if b.Type != "building_capital" {
			t.Errorf("building 0 type = %q", b.Type)
		}
		// "none" entries should be skipped
		if _, exists := gs.Districts[1]; exists {
			t.Error("district 1 (none) should have been skipped")
		}
	})

	// Leaders
	t.Run("leaders", func(t *testing.T) {
		if len(gs.Leaders) == 0 {
			t.Fatal("no leaders loaded")
		}
		// Leader 520093696 = Thalotha, scientist, country 0
		l, ok := gs.Leaders[520093696]
		if !ok {
			t.Fatal("leader 520093696 not found")
		}
		if l.Name.FullNames.Key != "HUM1_CHR_Thalotha" {
			t.Errorf("name = %q", l.Name.FullNames.Key)
		}
		if l.Class != "scientist" {
			t.Errorf("class = %q", l.Class)
		}
		if l.Country != 0 {
			t.Errorf("country = %d", l.Country)
		}
		if l.Age == 0 {
			t.Error("age is 0")
		}
		if len(l.Traits) == 0 {
			t.Error("no traits")
		}
		// "none" leader entries should be skipped
		if _, exists := gs.Leaders[654311425]; exists {
			t.Error("leader 654311425 (none) should have been skipped")
		}
	})
}
