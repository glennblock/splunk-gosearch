package goforit_command

import (
	"os"

	"./splunk/search"
)

type GoSearchCommand struct {
}

func (self *GoSearchCommand) HandleGetInfo(metadata commands.Metadata) commands.Metadata {
	return nil
}

func (self *GoSearchCommand) HandleExecute(metadata commands.Metadata, data commands.Data) commands.Data {
	return nil
}

func (self *GoSearchCommand) HandleResults(metadata commands.Metadata, data commands.Data) commands.Data {
	return nil
}

func main() {
	command := new(GoForItCommand)
	executor := commands.NewExecutor(os.Stdin, os.Stdout, new(GoSearchCommand))
}
