package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ashwins93/monkey/repl"
)

func main() {
	current, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Monkey\n", current.Username)

	repl.Start(os.Stdin, os.Stdout)
}
