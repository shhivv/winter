package lexer

import "winter/token"

type Lexer struct {
	input        string
	position     int  // points to current char
	readPosition int  // points to next char
	ch           byte // current char
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ascii for NUL, indicates that either no tokens have been created, or EOF has been reached
	} else {
		l.ch = l.input[l.readPosition] // string indexing produces ascii code (rune) instead of a string
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

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
		tok.Literal = "" // string(0)
		tok.Type = token.EOF
  default:
    if isLetter(l.ch){
      tok.Literal = l.readIdentifier()
      tok.Type  = token.LookupIdent(tok.Literal)
      return tok // early return because readIdentifier already advances position
    } else{
      tok = newToken(token.ILLEGAL, l.ch)
    }
	}
	l.readChar()
	return tok
}

// reads the next word until whitespace (can be a keyword or ident)
func (l *Lexer) readIdentifier() string{
  position := l.position
  for isLetter(l.ch){
    l.readChar()
  }

  return l.input[position:l.position]
}
func isLetter(ch byte) bool {
  return 'a' <= ch && ch >= 'z' || 'A' <= ch && ch >= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // consume first character into l.ch and init lexer
	return l
}
