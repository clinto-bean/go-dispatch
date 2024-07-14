package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// type config struct {
// // eventually will be able to cycle through all commands using arrow keys
// 	next *string
// 	prev *string
// }

func Repl() error {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println()
	fmt.Println("Welcome to go-dispatch! This lightweight worker will assist you with dispatching system commands and http handlers.")
	fmt.Println("A list of commands can be found below.")
	fmt.Println("If there are any concerns, please contact the maintainer.")
	fmt.Println()
	fmt.Println("To execute a command, enter 'do [name of command] [...args]'")
	fmt.Println()
	fmt.Print("> ")
	for {

		reader.Scan()
		input := reader.Text()
		words := strings.Split(input, " ")

		trigger := strings.ToLower(words[0])

		switch trigger {
		case "help":
			{
				GetAllCommands()
			}
		case "exit":
			{
				return errors.New("kill")
			}
		case "do":
			{
				if len(words) == 1 {
					//return error 5, no arguments passed
					return errors.New("not enough arguments passed")
				}
				command := strings.ToLower(words[1])
				c, err := GetCommand(command)
				if err != nil {
					return err
				}
				c.Func()
			}
		case "handle":
			{
				if len(words) == 1 {
					return errors.New("not enough arguments passed")

				}
				handler := strings.ToLower(words[1])
				fmt.Println(handler)

			}
		default:
			{
				fmt.Println("Input not recognized, type help to view available commands")
			}
		}

		fmt.Print("> ")
	}
}
