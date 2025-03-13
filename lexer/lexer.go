package lexer

import "monkey/token"

type Lexer struct {
	input string
	// POINTS to the current position corresponding to the `ch`
	position int // current position in input (points to current char)

	// ALWAYS points to the NEXT character in the `input`
	readPosition int // current reading position in input (after current char)

	ch byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// kind of misleading function name because it implies that we should read the next token
// but when initializing a lexer we start at the first character, assuming a token of `;`
// then calling `NextToken` returns that token of `;` which is the current token,
// that we are reading and not the NEXT token of the current one

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIndentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	// check if out of bounds
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// set new character from read position
		l.ch = l.input[l.readPosition]
	}
	// updated current position to next position in the input
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIndentifier() string {
	position := l.position
	// reads each letter of its corresponding asccii encoded value
	// and keeps moving the readPosition until it is not
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// checks the current character if it falls under the asccii encoded value
// for letters if so return true ch <-> ascii alphabet value
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

}
