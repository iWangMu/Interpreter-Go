package repl

import (
	"bufio"
	"fmt"
	"github.com/iWangMu/Interpreter-Go/lexer"
	"github.com/iWangMu/Interpreter-Go/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			if tok.Type == token.ILLEGAL {
				fmt.Println(fmt.Errorf("ILLEGAL input: %+v\n", tok))
			} else {
				fmt.Printf("%+v\n", tok)
			}
		}
	}
}
