package clausewitz

type tokenKind int

const (
	tokenEOF    tokenKind = iota
	tokenEquals           // =
	tokenLBrace           // {
	tokenRBrace           // }
	tokenIdent            // unquoted word, number, date, etc.
	tokenString           // "quoted string"
)

type token struct {
	kind  tokenKind
	value string
	line  int
}

type lexer struct {
	input     []byte
	pos       int
	line      int
	peeked    token
	hasPeeked bool
}

func newLexer(data []byte) *lexer {
	// Strip UTF-8 BOM
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		data = data[3:]
	}
	return &lexer{input: data, line: 1}
}

func (l *lexer) next() token {
	if l.hasPeeked {
		l.hasPeeked = false
		return l.peeked
	}
	return l.scan()
}

func (l *lexer) peek() token {
	if !l.hasPeeked {
		l.peeked = l.scan()
		l.hasPeeked = true
	}
	return l.peeked
}

func (l *lexer) scan() token {
	l.skipIgnored()
	if l.pos >= len(l.input) {
		return token{kind: tokenEOF, line: l.line}
	}
	switch l.input[l.pos] {
	case '=':
		l.pos++
		return token{kind: tokenEquals, line: l.line}
	case '{':
		l.pos++
		return token{kind: tokenLBrace, line: l.line}
	case '}':
		l.pos++
		return token{kind: tokenRBrace, line: l.line}
	case '"':
		return l.scanString()
	default:
		return l.scanIdent()
	}
}

func (l *lexer) skipIgnored() {
	for l.pos < len(l.input) {
		switch l.input[l.pos] {
		case ' ', '\t', '\r':
			l.pos++
		case '\n':
			l.pos++
			l.line++
		case '#':
			for l.pos < len(l.input) && l.input[l.pos] != '\n' {
				l.pos++
			}
		default:
			return
		}
	}
}

func (l *lexer) scanString() token {
	line := l.line
	l.pos++ // skip opening "
	var buf []byte
	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		if ch == '\\' && l.pos+1 < len(l.input) {
			next := l.input[l.pos+1]
			switch next {
			case '"', '\\':
				buf = append(buf, next)
			case 'n':
				buf = append(buf, '\n')
			case 't':
				buf = append(buf, '\t')
			default:
				buf = append(buf, ch, next)
			}
			l.pos += 2
			continue
		}
		if ch == '"' {
			l.pos++
			return token{kind: tokenString, value: string(buf), line: line}
		}
		if ch == '\n' {
			l.line++
		}
		buf = append(buf, ch)
		l.pos++
	}
	return token{kind: tokenString, value: string(buf), line: line}
}

func (l *lexer) scanIdent() token {
	start := l.pos
	line := l.line
	for l.pos < len(l.input) {
		switch l.input[l.pos] {
		case ' ', '\t', '\r', '\n', '=', '{', '}', '"', '#':
			return token{kind: tokenIdent, value: string(l.input[start:l.pos]), line: line}
		}
		l.pos++
	}
	return token{kind: tokenIdent, value: string(l.input[start:l.pos]), line: line}
}
