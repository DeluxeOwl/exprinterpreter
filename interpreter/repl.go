package interpreter

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">>"

func StartRepl(in io.Reader, ou io.Writer) {
	scanner := bufio.NewScanner(in)

	useLexer := false

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		if line == "use lexer" {
			useLexer = true
			fmt.Println("using lexer")
			continue
		}

		if line == "use parser" {
			useLexer = false
			fmt.Println("using parser")
			continue
		}

		if useLexer {
			l := NewLexer(line)
			for tok := l.NextToken(); tok.Type != EOFToken; tok = l.NextToken() {
				fmt.Printf("%+v\n", tok)
			}
		} else {
			l := NewLexer(line)
			p := NewParser(l)

			program := p.ParseProgram()
			fmt.Println(program)
		}
	}
}
