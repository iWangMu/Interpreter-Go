package token

// 词法单元类型
const (
	// 未知的词法单元/字符
	ILLEGAL TokenType = "ILLEGAL"
	// EndOfFile 文件结尾
	EOF TokenType = "EOF"

	// 标识符和字面量
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	// 运算符
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	// 分隔符
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"

	// 关键字
	FUNCTION TokenType = "FN"
	LET      TokenType = "LET"
)

type TokenType string

// 词法单元
type Token struct {
	Type    TokenType
	Literal string
}
