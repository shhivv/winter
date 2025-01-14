package main

import (
	"fmt"
	"os"
	"winter/repl"
)

func main() {
	fmt.Println("Winter 0.0.0")
	repl.Start(os.Stdin, os.Stdout)
}
