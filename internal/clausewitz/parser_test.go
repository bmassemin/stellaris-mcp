package clausewitz

import (
	"reflect"
	"testing"
)

// Test helpers for building expected AST values.
func sv(v string) Value         { return Value{Scalar: v} }
func ov(pairs ...Pair) Value    { return Value{Object: &Object{Pairs: pairs}} }
func lv(items ...Value) Value   { return Value{List: items} }
func kv(k string, v Value) Pair { return Pair{Key: k, Value: v} }

func TestParse_SimpleKeyValue(t *testing.T) {
	obj, err := Parse([]byte(`name="Earth" size=25`))
	if err != nil {
		t.Fatal(err)
	}
	want := &Object{Pairs: []Pair{
		kv("name", sv("Earth")),
		kv("size", sv("25")),
	}}
	if !reflect.DeepEqual(obj, want) {
		t.Errorf("\ngot:  %#v\nwant: %#v", obj, want)
	}
}

func TestParse_NestedObject(t *testing.T) {
	obj, err := Parse([]byte(`player={ name="Test" country=0 }`))
	if err != nil {
		t.Fatal(err)
	}
	want := &Object{Pairs: []Pair{
		kv("player", ov(
			kv("name", sv("Test")),
			kv("country", sv("0")),
		)),
	}}
	if !reflect.DeepEqual(obj, want) {
		t.Errorf("\ngot:  %#v\nwant: %#v", obj, want)
	}
}

func TestParse_ScalarList(t *testing.T) {
	obj, err := Parse([]byte(`ids={ 1 2 3 }`))
	if err != nil {
		t.Fatal(err)
	}
	want := &Object{Pairs: []Pair{
		kv("ids", lv(sv("1"), sv("2"), sv("3"))),
	}}
	if !reflect.DeepEqual(obj, want) {
		t.Errorf("\ngot:  %#v\nwant: %#v", obj, want)
	}
}

func TestParse_StringList(t *testing.T) {
	obj, err := Parse([]byte(`dlcs={ "Utopia" "Nemesis" }`))
	if err != nil {
		t.Fatal(err)
	}
	want := &Object{Pairs: []Pair{
		kv("dlcs", lv(sv("Utopia"), sv("Nemesis"))),
	}}
	if !reflect.DeepEqual(obj, want) {
		t.Errorf("\ngot:  %#v\nwant: %#v", obj, want)
	}
}

func TestParse_BlockList(t *testing.T) {
	input := `planets={
		{ name="Earth" size=25 }
		{ name="Mars" size=16 }
	}`
	obj, err := Parse([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	want := &Object{Pairs: []Pair{
		kv("planets", lv(
			ov(kv("name", sv("Earth")), kv("size", sv("25"))),
			ov(kv("name", sv("Mars")), kv("size", sv("16"))),
		)),
	}}
	if !reflect.DeepEqual(obj, want) {
		t.Errorf("\ngot:  %#v\nwant: %#v", obj, want)
	}
}

func TestParse_EmptyBlock(t *testing.T) {
	obj, err := Parse([]byte(`data={}`))
	if err != nil {
		t.Fatal(err)
	}
	want := &Object{Pairs: []Pair{
		kv("data", Value{Object: &Object{}}),
	}}
	if !reflect.DeepEqual(obj, want) {
		t.Errorf("\ngot:  %#v\nwant: %#v", obj, want)
	}
}

func TestParse_DuplicateKeys(t *testing.T) {
	obj, err := Parse([]byte(`trait=intelligent trait=adaptive`))
	if err != nil {
		t.Fatal(err)
	}
	want := &Object{Pairs: []Pair{
		kv("trait", sv("intelligent")),
		kv("trait", sv("adaptive")),
	}}
	if !reflect.DeepEqual(obj, want) {
		t.Errorf("\ngot:  %#v\nwant: %#v", obj, want)
	}
}

func TestParse_Comments(t *testing.T) {
	input := "# header\nname=Earth # inline\nsize=25"
	obj, err := Parse([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if len(obj.Pairs) != 2 {
		t.Fatalf("got %d pairs, want 2", len(obj.Pairs))
	}
	if obj.Pairs[0].Key != "name" || obj.Pairs[0].Value.Scalar != "Earth" {
		t.Errorf("pair[0] = %+v", obj.Pairs[0])
	}
	if obj.Pairs[1].Key != "size" || obj.Pairs[1].Value.Scalar != "25" {
		t.Errorf("pair[1] = %+v", obj.Pairs[1])
	}
}

func TestParse_DeepNesting(t *testing.T) {
	obj, err := Parse([]byte(`a={ b={ c={ d=1 } } }`))
	if err != nil {
		t.Fatal(err)
	}
	want := &Object{Pairs: []Pair{
		kv("a", ov(kv("b", ov(kv("c", ov(kv("d", sv("1")))))))),
	}}
	if !reflect.DeepEqual(obj, want) {
		t.Errorf("\ngot:  %#v\nwant: %#v", obj, want)
	}
}

func TestParse_NumericKeys(t *testing.T) {
	obj, err := Parse([]byte(`data={ 0=foo 1=bar }`))
	if err != nil {
		t.Fatal(err)
	}
	want := &Object{Pairs: []Pair{
		kv("data", ov(kv("0", sv("foo")), kv("1", sv("bar")))),
	}}
	if !reflect.DeepEqual(obj, want) {
		t.Errorf("\ngot:  %#v\nwant: %#v", obj, want)
	}
}

func TestParse_SingleElementList(t *testing.T) {
	obj, err := Parse([]byte(`ids={ 42 }`))
	if err != nil {
		t.Fatal(err)
	}
	want := &Object{Pairs: []Pair{
		kv("ids", lv(sv("42"))),
	}}
	if !reflect.DeepEqual(obj, want) {
		t.Errorf("\ngot:  %#v\nwant: %#v", obj, want)
	}
}

func TestParse_BooleanValues(t *testing.T) {
	obj, err := Parse([]byte(`active=yes deleted=no`))
	if err != nil {
		t.Fatal(err)
	}
	if obj.Pairs[0].Value.Scalar != "yes" {
		t.Errorf("expected 'yes', got %q", obj.Pairs[0].Value.Scalar)
	}
	if obj.Pairs[1].Value.Scalar != "no" {
		t.Errorf("expected 'no', got %q", obj.Pairs[1].Value.Scalar)
	}
}

func TestParse_ErrorMissingEquals(t *testing.T) {
	_, err := Parse([]byte(`key value`))
	if err == nil {
		t.Fatal("expected error for missing '='")
	}
}

func TestParse_ErrorUnterminatedBlock(t *testing.T) {
	_, err := Parse([]byte(`data={ key=val`))
	if err == nil {
		t.Fatal("expected error for unterminated block")
	}
}
