package main

import (
	"fmt"
	"os"
)

func main() {
	argsLength := len(os.Args)

	if argsLength == 1 {
		fmt.Println("No parameters were provided")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		if argsLength == 2 {
			if err := initialise(); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Repository initialised")
		} else {
			fmt.Println("init requires no parameters")
			return
		}

	case "commit":
		if argsLength == 3 {
			message := os.Args[2]

			id, err := createCommit(".", message)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Created commit", id)
		} else {
			fmt.Println("commit requires a message")
			return
		}
	default:
		fmt.Println("Unknown command", command)
		return
	}
}
