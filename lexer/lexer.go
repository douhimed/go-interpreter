package lexer

import "go/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = neWT
	}
}

func newToken(tt token.Token, ch byte) token.Token {
	return token.Token{Type: tt, Literal: string(ch)}
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= l.position {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
