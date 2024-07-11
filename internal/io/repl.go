package io

import (
	"bufio"
	"fmt"
	"os"
	// "github.com/clinto-bean/internal/commands"
)

type config struct {
	nextCommand *string
	prevCommand *string
}

func RunIORepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("")
		reader.Scan()

		cmd := reader.Text()
		err := validate(parse(cmd))
		if err != nil {
			// invalid
			panic("Invalid")
		}
	}
}

func parse(text string) string {
	return ""
}

func validate(text string) error {
	return nil
}

// func getCommands() map[string]commands.Command {
	
// }