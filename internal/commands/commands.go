package commands

import "fmt"

// declare commands here

type Command struct {
	Name        string
	Description string
	Func        func()
}

type List map[string]Command

var commandTest = Command{
	Name:        "test",
	Description: "display all commands",
	Func: func() {
		fmt.Println("test command executed properly")
	},
}

var Commands = List{
	"test": commandTest,
}

func Do(command string) bool {
	c, ok := Commands[command]
	if !ok {
		return false
	}
	c.Func()
	return true
}

func Get(name string) (Command, error) {
	c, ok := Commands[name]
	if !ok {
		return Command{}, fmt.Errorf("no command found for [%v]", name)
	}
	return c, nil
}

func GetAll() (List, bool) {
	if len(Commands) == 0 {
		return List{}, false
	}
	return Commands, true
}
