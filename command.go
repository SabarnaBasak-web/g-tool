package main

type Command struct {
	Name        string
	Description string
	Usage       string
}

var Commands = []Command{
	{
		Name:        "save",
		Description: "Add, commit, and push changes",
		Usage:       `gtool save "message"`,
	},
}
