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
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	EQ       TokenType = "=="
	NEQ      TokenType = "!="
	LT       TokenType = "<"
	LTE      TokenType = "<="
	GT       TokenType = ">"
	GTE      TokenType = ">="
	BANG     TokenType = "!"
	AND      TokenType = "&"
	OR       TokenType = "|"

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
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
)

type TokenType string

// 词法单元
type Token struct {
	Type    TokenType
	Literal string
}

// 关键词
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// 获取标识符类型
func LookupIdent(ident string) TokenType {
	tokenType, ok := keywords[ident]
	if ok {
		return tokenType
	} else {
		return IDENT
	}
}
