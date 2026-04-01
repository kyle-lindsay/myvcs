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
			}
		} else {
			fmt.Println("init requires no parameters")
			return
		}
	default:
		fmt.Println("Unknown command " + command)
		return
	}
}
