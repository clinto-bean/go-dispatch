package io

import (
	"errors"
	"fmt"

	commands "github.com/clinto-bean/go-dispatch/internal/commands"
)

func GetAllCommands() {
	list, ok := commands.GetAll()
	if !ok {
		fmt.Println("No commands found")
		return
	}
	fmt.Println("-")
	for _, c := range list {
		fmt.Println(c.Name+":", c.Description)
	}
	fmt.Println("-")
}

func GetCommand(name string) (commands.Command, error) {
	c, err := commands.Get(name)
	fmt.Println()
	fmt.Printf("Attempting to locate command {%v}\n", name)
	fmt.Println()
	if err != nil {
		msg := fmt.Sprintf("could not locate command [%v]", name)
		return commands.Command{}, errors.New(msg)
	}
	return c, nil
}
