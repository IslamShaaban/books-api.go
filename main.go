package main

import (
	"books-api/cmd"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	flags = flag.NewFlagSet("books-api", flag.ExitOnError)
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	commandExtension := ""
	if runtime.GOOS == "windows" {
		commandExtension = ".exe"
	}

	path := fmt.Sprintf("%s/tmp", directory)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	appState := make(chan string)
	closeApp := make(chan string)
	restarter := make(chan bool)

	go func() {
		switch args[0] {
		case "", "serve", "s":
			cmd.RunServer(directory, commandExtension, appState)
		}
	}()

	for {
		select {
		case <-c:
			fmt.Printf("Killing Server\n")
			go func() {
				appState <- "stop"
			}()
			time.Sleep(time.Second * 1)
			go func() {
				closeApp <- "done"
			}()
		case <-restarter:
			fmt.Printf("Restarting Server\n")
			go cmd.RunServer(directory, commandExtension, appState)
			time.Sleep(time.Second)
		case <-closeApp:
			fmt.Printf("Server Shutdown\n")
			return
		}
	}
}
