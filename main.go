package main

import (
	"os"

	"github.com/DeluxeOwl/pratt/interpreter"
)

func main() {
	interpreter.StartRepl(os.Stdin, os.Stdout)
}
