package commands

// declare commands here

type Command struct {
	name string
	level int
	path string
}

type List []Command 

var A = Command{
	name: "test",
	level: 0,
	path: "#",
}

var Commands = List{
	A,
}