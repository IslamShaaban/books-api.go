package cmd

import (
	"books-api/utils/command"
	"fmt"
	"strings"
)

func RunServer(directory string, commandExtension string, appState chan string) {
	command.RunCommand(fmt.Sprintf("go build -o %s/tmp/app%s %s/cmd/server/main.go", directory, commandExtension, directory), appState)
	command.RunCommand(fmt.Sprintf("%s/tmp/app%s", directory, commandExtension), appState)
}

func RunGenerator(directory string, commandExtension string, appState chan string, args []string) {
	command.RunCommand(fmt.Sprintf("go build -o %s/tmp/generator%s %s/cmd/generator/main.go", directory, commandExtension, directory), appState)
	command.RunCommand(fmt.Sprintf("%s/tmp/generator%s %s", directory, commandExtension, strings.Join(args[1:], " ")), appState)
}

func RunMigration(directory string, commandExtension string, appState chan string, args []string) {
	command.RunCommand(fmt.Sprintf("go build -o %s/tmp/migrate%s %s/cmd/migrations/main.go", directory, commandExtension, directory), appState)
	command.RunCommand(fmt.Sprintf("%s/tmp/migrate%s %s", directory, commandExtension, strings.Join(args[1:], " ")), appState)
}
