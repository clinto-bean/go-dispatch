package commands

// declare commands here

type Command struct {
	Name string
	Level int
	Path string
	Description string
	Status int
}

type List []Command 

var commandHelp = Command{
	Name: "help",
	Level: 0,
	Path: "#",
	Description: "display all commands",
	Status: 0,
}

var commandExit = Command{
	Name: "exit",
	Level: 0,
	Path: "",
	Description: "exit program",
	Status: 1,
}

var Commands = List{
	commandHelp, commandExit,
}