package clausewitz

import (
	"fmt"
	"reflect"
	"strconv"
)

// Unmarshal parses Clausewitz-formatted data and stores the result
// in the struct pointed to by v, similar to encoding/json.Unmarshal.
//
// Struct fields are matched to Clausewitz keys using the "clausewitz" tag,
// or the field name if no tag is set. Use `clausewitz:"-"` to skip a field.
//
// Type mapping:
//   - string, int*, uint*, float* <- scalar values
//   - bool <- "yes" / "no"
//   - struct <- nested { key=value } blocks
//   - []T <- value lists { a b c } or duplicate keys
//   - map[K]V <- { key=value } blocks as maps
func Unmarshal(data []byte, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("clausewitz: Unmarshal requires a non-nil pointer")
	}
	obj, err := Parse(data)
	if err != nil {
		return err
	}
	return decodeObject(obj, rv.Elem())
}

func decodeObject(obj *Object, rv reflect.Value) error {
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("clausewitz: cannot decode object into %s", rv.Type())
	}

	grouped := make(map[string][]Value)
	for _, p := range obj.Pairs {
		grouped[p.Key] = append(grouped[p.Key], p.Value)
	}

	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if !field.IsExported() {
			continue
		}
		key := fieldKey(field)
		if key == "-" {
			continue
		}
		vals, ok := grouped[key]
		if !ok {
			continue
		}
		fv := rv.Field(i)

		// Duplicate keys for a slice field -> collect all values
		if fv.Kind() == reflect.Slice && len(vals) > 1 {
			slice := reflect.MakeSlice(fv.Type(), len(vals), len(vals))
			for j, v := range vals {
				if err := decodeValue(v, slice.Index(j)); err != nil {
					return fmt.Errorf("clausewitz: field %q[%d]: %w", key, j, err)
				}
			}
			fv.Set(slice)
			continue
		}

		if err := decodeValue(vals[0], fv); err != nil {
			return fmt.Errorf("clausewitz: field %q: %w", key, err)
		}
	}
	return nil
}

func decodeValue(val Value, rv reflect.Value) error {
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.String:
		if !val.IsScalar() {
			return fmt.Errorf("cannot decode non-scalar into string")
		}
		rv.SetString(val.Scalar)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if !val.IsScalar() {
			return fmt.Errorf("cannot decode non-scalar into %s", rv.Type())
		}
		n, err := strconv.ParseInt(val.Scalar, 10, 64)
		if err != nil {
			return fmt.Errorf("cannot parse %q as %s: %w", val.Scalar, rv.Type(), err)
		}
		rv.SetInt(n)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if !val.IsScalar() {
			return fmt.Errorf("cannot decode non-scalar into %s", rv.Type())
		}
		n, err := strconv.ParseUint(val.Scalar, 10, 64)
		if err != nil {
			return fmt.Errorf("cannot parse %q as %s: %w", val.Scalar, rv.Type(), err)
		}
		rv.SetUint(n)

	case reflect.Float32, reflect.Float64:
		if !val.IsScalar() {
			return fmt.Errorf("cannot decode non-scalar into %s", rv.Type())
		}
		f, err := strconv.ParseFloat(val.Scalar, 64)
		if err != nil {
			return fmt.Errorf("cannot parse %q as %s: %w", val.Scalar, rv.Type(), err)
		}
		rv.SetFloat(f)

	case reflect.Bool:
		if !val.IsScalar() {
			return fmt.Errorf("cannot decode non-scalar into bool")
		}
		rv.SetBool(val.Scalar == "yes")

	case reflect.Struct:
		if !val.IsObject() {
			return fmt.Errorf("cannot decode non-object into %s", rv.Type())
		}
		return decodeObject(val.Object, rv)

	case reflect.Slice:
		return decodeSlice(val, rv)

	case reflect.Map:
		if !val.IsObject() {
			return fmt.Errorf("cannot decode non-object into %s", rv.Type())
		}
		return decodeMap(val.Object, rv)

	default:
		return fmt.Errorf("unsupported type %s", rv.Type())
	}
	return nil
}

func decodeSlice(val Value, rv reflect.Value) error {
	if val.IsList() {
		slice := reflect.MakeSlice(rv.Type(), len(val.List), len(val.List))
		for i, item := range val.List {
			if err := decodeValue(item, slice.Index(i)); err != nil {
				return fmt.Errorf("index %d: %w", i, err)
			}
		}
		rv.Set(slice)
		return nil
	}
	// Empty object {} -> empty slice (common in Stellaris saves)
	if val.IsObject() && len(val.Object.Pairs) == 0 {
		rv.Set(reflect.MakeSlice(rv.Type(), 0, 0))
		return nil
	}
	// Single value -> single-element slice
	slice := reflect.MakeSlice(rv.Type(), 1, 1)
	if err := decodeValue(val, slice.Index(0)); err != nil {
		return err
	}
	rv.Set(slice)
	return nil
}

func decodeMap(obj *Object, rv reflect.Value) error {
	if rv.IsNil() {
		rv.Set(reflect.MakeMap(rv.Type()))
	}
	keyType := rv.Type().Key()
	elemType := rv.Type().Elem()

	for _, pair := range obj.Pairs {
		// Skip "none" entries (common in Stellaris saves for empty/destroyed slots)
		if pair.Value.IsScalar() && pair.Value.Scalar == "none" {
			continue
		}
		key := reflect.New(keyType).Elem()
		if err := setFromString(key, pair.Key); err != nil {
			return fmt.Errorf("map key %q: %w", pair.Key, err)
		}
		val := reflect.New(elemType).Elem()
		if err := decodeValue(pair.Value, val); err != nil {
			return fmt.Errorf("map[%s]: %w", pair.Key, err)
		}
		rv.SetMapIndex(key, val)
	}
	return nil
}

func setFromString(rv reflect.Value, s string) error {
	switch rv.Kind() {
	case reflect.String:
		rv.SetString(s)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		rv.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		rv.SetUint(n)
	default:
		return fmt.Errorf("unsupported map key type %s", rv.Type())
	}
	return nil
}

func fieldKey(f reflect.StructField) string {
	if tag := f.Tag.Get("clausewitz"); tag != "" {
		return tag
	}
	return f.Name
}
