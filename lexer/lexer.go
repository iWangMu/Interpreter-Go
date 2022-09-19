package lexer

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

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

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
