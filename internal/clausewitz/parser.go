package clausewitz

import "fmt"

// Value represents a Clausewitz value: scalar, object, or list.
type Value struct {
	Scalar string
	Object *Object
	List   []Value
}

// IsScalar returns true if the value is a scalar (string, number, boolean).
func (v Value) IsScalar() bool { return v.Object == nil && v.List == nil }

// IsObject returns true if the value is a key-value block.
func (v Value) IsObject() bool { return v.Object != nil }

// IsList returns true if the value is a list of values.
func (v Value) IsList() bool { return v.List != nil }

// Object is an ordered collection of key-value pairs.
// Keys may be duplicated.
type Object struct {
	Pairs []Pair
}

// Pair is a key-value entry in an Object.
type Pair struct {
	Key   string
	Value Value
}

// Parse parses Clausewitz-formatted data into an Object.
func Parse(data []byte) (*Object, error) {
	p := &parser{lex: newLexer(data)}
	return p.parseDocument()
}

type parser struct {
	lex *lexer
}

func (p *parser) parseDocument() (*Object, error) {
	var pairs []Pair
	for p.lex.peek().kind != tokenEOF {
		pair, err := p.parsePair()
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}
	return &Object{Pairs: pairs}, nil
}

func (p *parser) parsePair() (Pair, error) {
	key := p.lex.next()
	if key.kind != tokenIdent && key.kind != tokenString {
		return Pair{}, fmt.Errorf("line %d: expected key, got %v", key.line, key.kind)
	}
	eq := p.lex.next()
	if eq.kind != tokenEquals {
		return Pair{}, fmt.Errorf("line %d: expected '=' after key %q", eq.line, key.value)
	}
	val, err := p.parseValue()
	if err != nil {
		return Pair{}, err
	}
	return Pair{Key: key.value, Value: val}, nil
}

func (p *parser) parseValue() (Value, error) {
	tok := p.lex.peek()
	switch tok.kind {
	case tokenLBrace:
		return p.parseBlock()
	case tokenIdent, tokenString:
		p.lex.next()
		return Value{Scalar: tok.value}, nil
	default:
		return Value{}, fmt.Errorf("line %d: expected value", tok.line)
	}
}

// parseBlock parses a brace-delimited block.
// It disambiguates between key-value objects and value lists
// by peeking at the first two tokens inside the braces.
func (p *parser) parseBlock() (Value, error) {
	p.lex.next() // consume {

	first := p.lex.peek()

	if first.kind == tokenRBrace {
		p.lex.next()
		return Value{Object: &Object{}}, nil
	}

	// Starts with { -> list of blocks/values
	if first.kind == tokenLBrace {
		return p.parseListUntilClose()
	}

	// Consume first token and peek at second to determine block type
	p.lex.next()
	second := p.lex.peek()

	if second.kind == tokenEquals {
		return p.parsePairsFrom(first)
	}
	return p.parseValuesFrom(first)
}

// parsePairsFrom finishes parsing an object block where first key is already consumed.
func (p *parser) parsePairsFrom(firstKey token) (Value, error) {
	p.lex.next() // consume =
	val, err := p.parseValue()
	if err != nil {
		return Value{}, err
	}
	pairs := []Pair{{Key: firstKey.value, Value: val}}
	for p.lex.peek().kind != tokenRBrace {
		if p.lex.peek().kind == tokenEOF {
			return Value{}, fmt.Errorf("line %d: unterminated block", p.lex.peek().line)
		}
		pair, err := p.parsePair()
		if err != nil {
			return Value{}, err
		}
		pairs = append(pairs, pair)
	}
	p.lex.next() // consume }
	return Value{Object: &Object{Pairs: pairs}}, nil
}

// parseValuesFrom finishes parsing a value list where first item is already consumed.
func (p *parser) parseValuesFrom(first token) (Value, error) {
	items := []Value{{Scalar: first.value}}
	for p.lex.peek().kind != tokenRBrace {
		if p.lex.peek().kind == tokenEOF {
			return Value{}, fmt.Errorf("line %d: unterminated list", p.lex.peek().line)
		}
		val, err := p.parseValue()
		if err != nil {
			return Value{}, err
		}
		items = append(items, val)
	}
	p.lex.next() // consume }
	return Value{List: items}, nil
}

// parseListUntilClose parses values until a closing brace.
func (p *parser) parseListUntilClose() (Value, error) {
	var items []Value
	for p.lex.peek().kind != tokenRBrace {
		if p.lex.peek().kind == tokenEOF {
			return Value{}, fmt.Errorf("line %d: unterminated list", p.lex.peek().line)
		}
		val, err := p.parseValue()
		if err != nil {
			return Value{}, err
		}
		items = append(items, val)
	}
	p.lex.next() // consume }
	return Value{List: items}, nil
}
