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

		for _, tok := range l.Lex() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
