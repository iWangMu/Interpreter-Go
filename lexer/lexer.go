package lexer

import "github.com/iWangMu/Interpreter-Go/token"

type Lexer struct {
	// 输入字符串
	input string
	// 输入字符串中的当前位置（指向当前字符）
	position int
	// 输入字符串中的当前读取位置（指向当前字符之后的一个字符）
	readPosition int
	// 当前正在查看的字符
	ch byte
}

// 初始化
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

// 提取词法单元
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}
	l.readChar()
	return tok
}

// 读取input的下一个字符，只支持ASCII字符集
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// ASCII编码表示NUL字符，表示尚未读取任何内容或文件结尾
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {

	return token.Token{Type: tokenType, Literal: string(ch)}
}
