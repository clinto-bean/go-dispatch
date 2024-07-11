package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/clinto-bean/go-dispatch/internal/commands"
)

// type config struct {
// 	next *string
// 	prev *string
// }

func RunIORepl() int {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println()
	fmt.Println("To execute a command, enter 'do [name of command] [...args]'")
	fmt.Println()
	fmt.Print("> ")
	for {
		reader.Scan()
		cmd := reader.Text()
		words := strings.Split(cmd, " ")
		err := parse(words)
		if err > 0 {
			return err
		}
		if err == 0 {
			fmt.Println()
			for _, c := range getCommands() {
				fmt.Printf("%v: %v\n", c.Name, c.Description)
			}
			fmt.Println()
		}
		fmt.Print("> ")
	}
}


func parse(text []string) (int) {
	switch text[0] {
		case "help": return 0
		case "exit": return 1
		default: return 3
	}
}

func getCommands() []commands.Command {
	return commands.Commands
}

// func getHandlers() map[string]handlers.Handler {
// 	return nil
// }