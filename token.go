package migo // import "github.com/dmarottoli/migo"

import "fmt"

// Tokens for use with lexer and parser.

// Tok is a lexical token.
type Tok int

// Token is a token with metadata.
type Token interface {
	Tok() Tok
	StartPos() TokenPos
	EndPos() TokenPos
}

// ConstToken is a normal constant token.
type ConstToken struct {
	t          Tok
	start, end TokenPos
}

// Tok returns the token id.
func (t *ConstToken) Tok() Tok {
	return t.t
}

// StartPos returns starting position of token.
func (t *ConstToken) StartPos() TokenPos {
	return t.start
}

// EndPos returns ending position of token.
func (t *ConstToken) EndPos() TokenPos {
	return t.end
}

// IdentToken is a token with string value (Ident).
type IdentToken struct {
	str        string
	start, end TokenPos
}

// Tok returns IDENT.
func (*IdentToken) Tok() Tok {
	return IDENT
}

// StartPos returns starting position of token.
func (t *IdentToken) StartPos() TokenPos {
	return t.start
}

// EndPos returns ending position of token.
func (t *IdentToken) EndPos() TokenPos {
	return t.end
}

// DigitsToken is a token with numeric value (Digits).
type DigitsToken struct {
	num        int
	start, end TokenPos
}

// Tok returns DIGITS.
func (t *DigitsToken) Tok() Tok {
	return DIGITS
}

// StartPos returns starting position of token.
func (t *DigitsToken) StartPos() TokenPos {
	return t.start
}

// EndPos returns ending position of token.
func (t *DigitsToken) EndPos() TokenPos {
	return t.end
}

const (
	// ILLEGAL is a special token for errors.
	ILLEGAL Tok = iota
)

var eof = rune(0)

// TokenPos is a pair of coordinate to identify start of token.
type TokenPos struct {
	Char  int
	Lines []int
}

func (p TokenPos) String() string {
	return fmt.Sprintf("%d:%d", len(p.Lines)+1, p.Char)
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isAlphaNum(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ('0' <= ch && ch <= '9')
}

func isNum(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isIdent(ch rune) bool {
	return isAlphaNum(ch) || ch == '_' || ch == '.' || ch == '#' || ch == '/'
}
