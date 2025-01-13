package interpreter

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">>"

func StartRepl(in io.Reader, ou io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := NewLexer(line)

		for tok := l.NextToken(); tok.Type != EOFToken; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
