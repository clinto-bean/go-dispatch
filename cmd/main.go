package main

import (
	"fmt"
	"os"

	io "github.com/clinto-bean/go-dispatch/internal/io"
)

func main() {
	fmt.Println()
	fmt.Println("Welcome to go-dispatch! This lightweight worker will assist you with dispatching system commands and http handlers.")
	fmt.Println("A list of commands can be found below.")
	fmt.Println("If there are any concerns, please contact the maintainer.")
	err := io.RunIORepl()
	switch err {
	case 1:
		fmt.Println("EXIT: Exiting")
	default:
		fmt.Printf("Uncaught exception, code %v\n", err)
	}
	os.Exit(err)
}
