package main

import (
	"os"

	"github.com/DeluxeOwl/exprinterpreter/interpreter"
)

func main() {
	interpreter.StartRepl(os.Stdin, os.Stdout)
}
