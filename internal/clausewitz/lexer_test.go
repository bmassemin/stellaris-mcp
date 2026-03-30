package clausewitz

import "testing"

func collectTokens(input string) []token {
	l := newLexer([]byte(input))
	var tokens []token
	for {
		tok := l.next()
		tokens = append(tokens, tok)
		if tok.kind == tokenEOF {
			break
		}
	}
	return tokens
}

func TestLexer_BasicTokens(t *testing.T) {
	tokens := collectTokens(`key = { }`)
	expected := []tokenKind{tokenIdent, tokenEquals, tokenLBrace, tokenRBrace, tokenEOF}
	if len(tokens) != len(expected) {
		t.Fatalf("got %d tokens, want %d", len(tokens), len(expected))
	}
	for i, tok := range tokens {
		if tok.kind != expected[i] {
			t.Errorf("token[%d]: got kind %d, want %d", i, tok.kind, expected[i])
		}
	}
	if tokens[0].value != "key" {
		t.Errorf("token[0].value = %q, want %q", tokens[0].value, "key")
	}
}

func TestLexer_Identifiers(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"hello", "hello"},
		{"42", "42"},
		{"-3.14", "-3.14"},
		{"2200.01.01", "2200.01.01"},
		{"star_class_a", "star_class_a"},
		{"@base_value", "@base_value"},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			l := newLexer([]byte(tc.input))
			tok := l.next()
			if tok.kind != tokenIdent {
				t.Fatalf("got kind %d, want tokenIdent", tok.kind)
			}
			if tok.value != tc.want {
				t.Errorf("got %q, want %q", tok.value, tc.want)
			}
		})
	}
}

func TestLexer_Strings(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"simple", `"hello"`, "hello"},
		{"spaces", `"hello world"`, "hello world"},
		{"escaped_quote", `"say \"hi\""`, `say "hi"`},
		{"escaped_backslash", `"path\\to"`, `path\to`},
		{"escaped_newline", `"line1\nline2"`, "line1\nline2"},
		{"empty", `""`, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := newLexer([]byte(tc.input))
			tok := l.next()
			if tok.kind != tokenString {
				t.Fatalf("got kind %d, want tokenString", tok.kind)
			}
			if tok.value != tc.want {
				t.Errorf("got %q, want %q", tok.value, tc.want)
			}
		})
	}
}

func TestLexer_Comments(t *testing.T) {
	tokens := collectTokens("key=value # comment\nother=thing")
	kinds := []tokenKind{tokenIdent, tokenEquals, tokenIdent, tokenIdent, tokenEquals, tokenIdent, tokenEOF}
	if len(tokens) != len(kinds) {
		t.Fatalf("got %d tokens, want %d", len(tokens), len(kinds))
	}
	for i, tok := range tokens {
		if tok.kind != kinds[i] {
			t.Errorf("token[%d]: got kind %d, want %d", i, tok.kind, kinds[i])
		}
	}
}

func TestLexer_CommentOnly(t *testing.T) {
	tokens := collectTokens("# just a comment")
	if len(tokens) != 1 || tokens[0].kind != tokenEOF {
		t.Errorf("expected only EOF, got %d tokens", len(tokens))
	}
}

func TestLexer_BOM(t *testing.T) {
	input := "\xEF\xBB\xBFkey=value"
	l := newLexer([]byte(input))
	tok := l.next()
	if tok.kind != tokenIdent || tok.value != "key" {
		t.Errorf("got {%d, %q}, want {tokenIdent, \"key\"}", tok.kind, tok.value)
	}
}

func TestLexer_LineNumbers(t *testing.T) {
	tokens := collectTokens("a=1\nb=2\nc=3")
	if tokens[0].line != 1 {
		t.Errorf("first token line = %d, want 1", tokens[0].line)
	}
	// b is on line 2 (tokens: a=1 b=2 c=3 -> indices 0,1,2,3,4,5,6,7,8)
	if tokens[3].line != 2 {
		t.Errorf("'b' token line = %d, want 2", tokens[3].line)
	}
	if tokens[6].line != 3 {
		t.Errorf("'c' token line = %d, want 3", tokens[6].line)
	}
}

func TestLexer_Empty(t *testing.T) {
	tokens := collectTokens("")
	if len(tokens) != 1 || tokens[0].kind != tokenEOF {
		t.Errorf("expected only EOF for empty input")
	}
}

func TestLexer_Peek(t *testing.T) {
	l := newLexer([]byte("a=b"))
	// Peek should return the same token multiple times
	tok1 := l.peek()
	tok2 := l.peek()
	if tok1 != tok2 {
		t.Errorf("peek not idempotent: %v vs %v", tok1, tok2)
	}
	// Next should return the peeked token
	tok3 := l.next()
	if tok3 != tok1 {
		t.Errorf("next after peek: got %v, want %v", tok3, tok1)
	}
	// Next call should advance
	tok4 := l.next()
	if tok4.kind != tokenEquals {
		t.Errorf("expected tokenEquals, got %d", tok4.kind)
	}
}
