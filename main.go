package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/aaylward/goterp/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s. This is a toy interpreter. Good luck!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
