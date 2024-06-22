package command

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"math/rand"
	"strings"
	"time"
)

func RunCommand(command string, appState chan string) {
	id := rand.Intn(100000)

	pieces := strings.Split(command, " ")
	c := cmd.NewCmd(pieces[0], pieces[1:]...)

	lastStdout := 0
	lastStderr := 0
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			status := c.Status()
			n := len(status.Stdout)
			if n > 0 {
				for i := lastStdout; i < n; i++ {
					fmt.Println(status.Stdout[i])
				}
				lastStdout = n
			}
			n = len(status.Stderr)
			if n > 0 {
				for i := lastStderr; i < n; i++ {
					fmt.Println(status.Stderr[i])
				}
				lastStderr = n
			}
		}
	}()

	statusChan := c.Start()
	commandChan := make(chan string)

	go func() {
		complete := false
		for {
			if complete {
				fmt.Printf("Command complete complete loop: %s %d \n", command, id)
				break
			}
			select {
			case state := <-appState:
				if state == "restart" || state == "stop" {
					err := c.Stop()
					if err != nil {
						fmt.Printf("Error Killing command: %s \n", err.Error())
					}
				}
			case <-commandChan:
				fmt.Printf("Command complete kill loop: %s %d \n", command, id)
				complete = true
			}
		}
	}()

	// this waits until the command is done or killed
	<-statusChan

	commandChan <- "done"

	// dump the remaining contents in stdout and stderr
	status := c.Status()
	n := len(status.Stdout)
	if n > 0 {
		for i := lastStdout; i < n; i++ {
			fmt.Println(status.Stdout[i])
		}
		lastStdout = n
	}
	n = len(status.Stderr)
	if n > 0 {
		for i := lastStderr; i < n; i++ {
			fmt.Println(status.Stderr[i])
		}
		lastStderr = n
	}
}
