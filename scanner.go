package migo // import "github.com/dmarottoli/migo"

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

// Scanner is a lexical scanner.
type Scanner struct {
	r   *bufio.Reader
	pos TokenPos
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r), pos: TokenPos{Char: 0, Lines: []int{}}}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if reached the end or error occurs.
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	if ch == '\n' {
		s.pos.Lines = append(s.pos.Lines, s.pos.Char)
		s.pos.Char = 0
	} else {
		s.pos.Char++
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
	if s.pos.Char == 0 {
		s.pos.Char = s.pos.Lines[len(s.pos.Lines)-1]
		s.pos.Lines = s.pos.Lines[:len(s.pos.Lines)-1]
	} else {
		s.pos.Char--
	}
}

// Scan returns the next token and parsed value.
func (s *Scanner) Scan() Token {
	var startPos, endPos TokenPos
	ch := s.read()

	if isWhitespace(ch) {
		s.skipWhitespace()
		ch = s.read()
	}
	if isIdent(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Track token positions.
	startPos = s.pos
	defer func() { endPos = s.pos }()

	switch ch {
	case eof:
		return &ConstToken{t: 0, start: startPos, end: endPos}
	case ':':
		return &ConstToken{t: COLON, start: startPos, end: endPos}
	case ';':
		return &ConstToken{t: SEMICOLON, start: startPos, end: endPos}
	case ',':
		return &ConstToken{t: COMMA, start: startPos, end: endPos}
	case '(':
		return &ConstToken{t: LPAREN, start: startPos, end: endPos}
	case ')':
		return &ConstToken{t: RPAREN, start: startPos, end: endPos}
	case '=':
		return &ConstToken{t: EQ, start: startPos, end: endPos}
	case '-':
		if ch2 := s.read(); ch2 == '-' {
			s.unread()
			s.unread()
			s.skipComment()
			return s.Scan()
		}
	}
	return &ConstToken{t: ILLEGAL, start: startPos, end: endPos}
}

func (s *Scanner) scanIdent() Token {
	var startPos, endPos TokenPos
	var buf bytes.Buffer

	startPos = s.pos
	defer func() { endPos = s.pos }()

	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isIdent(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	switch buf.String() {
	case "def":
		return &ConstToken{t: DEF, start: startPos, end: endPos}
	case "call":
		return &ConstToken{t: CALL, start: startPos, end: endPos}
	case "spawn":
		return &ConstToken{t: SPAWN, start: startPos, end: endPos}
	case "case":
		return &ConstToken{t: CASE, start: startPos, end: endPos}
	case "close":
		return &ConstToken{t: CLOSE, start: startPos, end: endPos}
	case "else":
		return &ConstToken{t: ELSE, start: startPos, end: endPos}
	case "endif":
		return &ConstToken{t: ENDIF, start: startPos, end: endPos}
	case "endselect":
		return &ConstToken{t: ENDSELECT, start: startPos, end: endPos}
	case "if":
		return &ConstToken{t: IF, start: startPos, end: endPos}
	case "let":
		return &ConstToken{t: LET, start: startPos, end: endPos}
	case "newchan":
		return &ConstToken{t: NEWCHAN, start: startPos, end: endPos}
	case "select":
		return &ConstToken{t: SELECT, start: startPos, end: endPos}
	case "send":
		return &ConstToken{t: SEND, start: startPos, end: endPos}
	case "recv":
		return &ConstToken{t: RECV, start: startPos, end: endPos}
	case "tau":
		return &ConstToken{t: TAU, start: startPos, end: endPos}
	}

	if i, err := strconv.Atoi(buf.String()); err == nil {
		return &DigitsToken{num: i, start: startPos, end: endPos}
	}
	return &IdentToken{str: buf.String(), start: startPos, end: endPos}
}

func (s *Scanner) skipComment() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == '\n' {
			break
		}
	}
}

func (s *Scanner) skipWhitespace() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		}
	}
}
