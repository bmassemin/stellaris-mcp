package clausewitz

import (
	"os"
	"reflect"
	"testing"
)

func TestUnmarshal_Scalars(t *testing.T) {
	input := `name="Earth" size=25 distance=1.5 habitable=yes barren=no`
	type Planet struct {
		Name      string  `clausewitz:"name"`
		Size      int     `clausewitz:"size"`
		Distance  float64 `clausewitz:"distance"`
		Habitable bool    `clausewitz:"habitable"`
		Barren    bool    `clausewitz:"barren"`
	}
	var got Planet
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	want := Planet{Name: "Earth", Size: 25, Distance: 1.5, Habitable: true, Barren: false}
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestUnmarshal_IntTypes(t *testing.T) {
	input := `a=42 b=-10 c=255`
	type Data struct {
		A int32 `clausewitz:"a"`
		B int64 `clausewitz:"b"`
		C uint8 `clausewitz:"c"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	if got.A != 42 || got.B != -10 || got.C != 255 {
		t.Errorf("got %+v", got)
	}
}

func TestUnmarshal_NestedStruct(t *testing.T) {
	input := `player={ name="Test" country=10 }`
	type Player struct {
		Name    string `clausewitz:"name"`
		Country int    `clausewitz:"country"`
	}
	type Save struct {
		Player Player `clausewitz:"player"`
	}
	var got Save
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	if got.Player.Name != "Test" || got.Player.Country != 10 {
		t.Errorf("got %+v", got.Player)
	}
}

func TestUnmarshal_SliceFromList(t *testing.T) {
	input := `ids={ 10 20 30 }`
	type Data struct {
		IDs []int `clausewitz:"ids"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	want := []int{10, 20, 30}
	if !reflect.DeepEqual(got.IDs, want) {
		t.Errorf("got %v, want %v", got.IDs, want)
	}
}

func TestUnmarshal_StringSliceFromList(t *testing.T) {
	input := `dlcs={ "Utopia" "Nemesis" }`
	type Data struct {
		DLCs []string `clausewitz:"dlcs"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	want := []string{"Utopia", "Nemesis"}
	if !reflect.DeepEqual(got.DLCs, want) {
		t.Errorf("got %v, want %v", got.DLCs, want)
	}
}

func TestUnmarshal_SliceFromDuplicateKeys(t *testing.T) {
	input := `trait=industrious trait=intelligent`
	type Data struct {
		Trait []string `clausewitz:"trait"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	want := []string{"industrious", "intelligent"}
	if !reflect.DeepEqual(got.Trait, want) {
		t.Errorf("got %v, want %v", got.Trait, want)
	}
}

func TestUnmarshal_SliceOfStructs(t *testing.T) {
	input := `planets={
		{ name="Earth" size=25 }
		{ name="Mars" size=16 }
	}`
	type Planet struct {
		Name string `clausewitz:"name"`
		Size int    `clausewitz:"size"`
	}
	type Data struct {
		Planets []Planet `clausewitz:"planets"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	want := []Planet{{Name: "Earth", Size: 25}, {Name: "Mars", Size: 16}}
	if !reflect.DeepEqual(got.Planets, want) {
		t.Errorf("got %+v, want %+v", got.Planets, want)
	}
}

func TestUnmarshal_EmptyBlockAsSlice(t *testing.T) {
	input := `ids={}`
	type Data struct {
		IDs []int `clausewitz:"ids"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	if got.IDs == nil || len(got.IDs) != 0 {
		t.Errorf("expected non-nil empty slice, got %v", got.IDs)
	}
}

func TestUnmarshal_EmptyBlockAsStruct(t *testing.T) {
	input := `data={}`
	type Inner struct {
		Name string `clausewitz:"name"`
	}
	type Data struct {
		Data Inner `clausewitz:"data"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	if got.Data.Name != "" {
		t.Errorf("expected zero struct, got %+v", got.Data)
	}
}

func TestUnmarshal_Map(t *testing.T) {
	input := `stats={ economy=100 military=50 }`
	type Data struct {
		Stats map[string]int `clausewitz:"stats"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	want := map[string]int{"economy": 100, "military": 50}
	if !reflect.DeepEqual(got.Stats, want) {
		t.Errorf("got %v, want %v", got.Stats, want)
	}
}

func TestUnmarshal_MapIntKeys(t *testing.T) {
	input := `data={ 0=hello 1=world }`
	type Data struct {
		Data map[int]string `clausewitz:"data"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	want := map[int]string{0: "hello", 1: "world"}
	if !reflect.DeepEqual(got.Data, want) {
		t.Errorf("got %v, want %v", got.Data, want)
	}
}

func TestUnmarshal_MapOfStructs(t *testing.T) {
	input := `planets={ earth={ size=25 } mars={ size=16 } }`
	type Planet struct {
		Size int `clausewitz:"size"`
	}
	type Data struct {
		Planets map[string]Planet `clausewitz:"planets"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	if got.Planets["earth"].Size != 25 || got.Planets["mars"].Size != 16 {
		t.Errorf("got %+v", got.Planets)
	}
}

func TestUnmarshal_Pointer(t *testing.T) {
	input := `name="Earth" size=25`
	type Data struct {
		Name *string `clausewitz:"name"`
		Size *int    `clausewitz:"size"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	if got.Name == nil || *got.Name != "Earth" {
		t.Errorf("Name = %v", got.Name)
	}
	if got.Size == nil || *got.Size != 25 {
		t.Errorf("Size = %v", got.Size)
	}
}

func TestUnmarshal_NilPointerUntouched(t *testing.T) {
	input := `name="Earth"`
	type Data struct {
		Name string `clausewitz:"name"`
		Size *int   `clausewitz:"size"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	if got.Size != nil {
		t.Errorf("Size should be nil, got %v", got.Size)
	}
}

func TestUnmarshal_TagOverride(t *testing.T) {
	input := `custom_name="test"`
	type Data struct {
		Name string `clausewitz:"custom_name"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	if got.Name != "test" {
		t.Errorf("got %q, want %q", got.Name, "test")
	}
}

func TestUnmarshal_TagSkip(t *testing.T) {
	input := `name="Earth" internal=secret`
	type Data struct {
		Name     string `clausewitz:"name"`
		Internal string `clausewitz:"-"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	if got.Internal != "" {
		t.Errorf("Internal should be empty, got %q", got.Internal)
	}
}

func TestUnmarshal_MissingFields(t *testing.T) {
	input := `name="Earth"`
	type Data struct {
		Name string `clausewitz:"name"`
		Size int    `clausewitz:"size"`
	}
	var got Data
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}
	if got.Name != "Earth" || got.Size != 0 {
		t.Errorf("got %+v", got)
	}
}

func TestUnmarshal_Stellaris(t *testing.T) {
	input := `
version="3.6.1"
name="United Nations of Earth"
date="2200.03.01"
player={
	name="Player"
	country=0
}
required_dlcs={
	"Synthetic Dawn Story Pack"
	"Utopia"
}
species={
	{
		name="Human"
		plural="Humans"
		traits={
			trait=intelligent
			trait=adaptive
		}
	}
	{
		name="Blorg"
		plural="Blorg"
		traits={
			trait=charismatic
		}
	}
}
`
	type Traits struct {
		Trait []string `clausewitz:"trait"`
	}
	type Species struct {
		Name   string `clausewitz:"name"`
		Plural string `clausewitz:"plural"`
		Traits Traits `clausewitz:"traits"`
	}
	type Player struct {
		Name    string `clausewitz:"name"`
		Country int    `clausewitz:"country"`
	}
	type SaveGame struct {
		Version      string    `clausewitz:"version"`
		Name         string    `clausewitz:"name"`
		Date         string    `clausewitz:"date"`
		Player       Player    `clausewitz:"player"`
		RequiredDLCs []string  `clausewitz:"required_dlcs"`
		Species      []Species `clausewitz:"species"`
	}

	var got SaveGame
	if err := Unmarshal([]byte(input), &got); err != nil {
		t.Fatal(err)
	}

	if got.Version != "3.6.1" {
		t.Errorf("Version = %q", got.Version)
	}
	if got.Name != "United Nations of Earth" {
		t.Errorf("Name = %q", got.Name)
	}
	if got.Date != "2200.03.01" {
		t.Errorf("Date = %q", got.Date)
	}
	if got.Player.Name != "Player" || got.Player.Country != 0 {
		t.Errorf("Player = %+v", got.Player)
	}
	wantDLCs := []string{"Synthetic Dawn Story Pack", "Utopia"}
	if !reflect.DeepEqual(got.RequiredDLCs, wantDLCs) {
		t.Errorf("RequiredDLCs = %v", got.RequiredDLCs)
	}
	if len(got.Species) != 2 {
		t.Fatalf("Species count = %d, want 2", len(got.Species))
	}
	if got.Species[0].Name != "Human" || got.Species[0].Plural != "Humans" {
		t.Errorf("Species[0] = %+v", got.Species[0])
	}
	wantTraits0 := []string{"intelligent", "adaptive"}
	if !reflect.DeepEqual(got.Species[0].Traits.Trait, wantTraits0) {
		t.Errorf("Species[0].Traits = %v", got.Species[0].Traits.Trait)
	}
	if got.Species[1].Name != "Blorg" {
		t.Errorf("Species[1].Name = %q", got.Species[1].Name)
	}
	wantTraits1 := []string{"charismatic"}
	if !reflect.DeepEqual(got.Species[1].Traits.Trait, wantTraits1) {
		t.Errorf("Species[1].Traits = %v", got.Species[1].Traits.Trait)
	}
}

func TestParse_GamestateFixture(t *testing.T) {
	data, err := os.ReadFile("testdata/gamestate")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}
	obj, err := Parse(data)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if len(obj.Pairs) == 0 {
		t.Fatal("parsed zero top-level pairs")
	}
}

func TestUnmarshal_GamestateFixture(t *testing.T) {
	data, err := os.ReadFile("testdata/gamestate")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	type GameState struct {
		Version      string   `clausewitz:"version"`
		Name         string   `clausewitz:"name"`
		Date         string   `clausewitz:"date"`
		RequiredDLCs []string `clausewitz:"required_dlcs"`
	}

	var gs GameState
	if err := Unmarshal(data, &gs); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if gs.Version != "Cetus v4.3.2" {
		t.Errorf("Version = %q", gs.Version)
	}
	if gs.Name != "mp_Bebakian League" {
		t.Errorf("Name = %q", gs.Name)
	}
	if gs.Date != "2200.07.01" {
		t.Errorf("Date = %q", gs.Date)
	}
	if len(gs.RequiredDLCs) == 0 {
		t.Error("RequiredDLCs is empty")
	}
}

func TestUnmarshal_ErrorNonPointer(t *testing.T) {
	type Data struct{}
	var d Data
	err := Unmarshal([]byte(`key=val`), d)
	if err == nil {
		t.Fatal("expected error for non-pointer")
	}
}

func TestUnmarshal_ErrorNilPointer(t *testing.T) {
	err := Unmarshal([]byte(`key=val`), (*struct{ X string })(nil))
	if err == nil {
		t.Fatal("expected error for nil pointer")
	}
}
