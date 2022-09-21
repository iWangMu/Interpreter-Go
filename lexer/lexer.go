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
	l := &Lexer{input: input, ch: ' '}
	return l
}

// 提取词法单元
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	// 跳过空白字符
	l.skipWhitespace()
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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// 此时l中的ch已经是指向下一个字符了
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
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

// 获取完整的字符串标识
func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) readNumber() string {
	startPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 只支持中英文大小写字符和下划线
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 数字
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
