package cmd

import (
	"books-api/utils/command"
	"fmt"
)

func RunServer(directory string, commandExtension string, appState chan string) {
	command.RunCommand(fmt.Sprintf("go build -o %s/tmp/app%s %s/cmd/server/main.go", directory, commandExtension, directory), appState)
	command.RunCommand(fmt.Sprintf("%s/tmp/app%s", directory, commandExtension), appState)
}