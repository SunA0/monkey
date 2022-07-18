package main

import (
	"fmt"
	"monkey/repl"
	"os"
	usr "os/user"
)

func main() {
	user, err := usr.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
