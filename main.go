package main

import (
	"fmt"
	"os"
)

func initialise() error {

	_, err := os.Stat(".myvcs")
	if err == nil {
		return fmt.Errorf("repository already exists in this directory")
	}
	if !os.IsNotExist(err) {
		return err
	}

	if err := os.MkdirAll(".myvcs", 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(".myvcs/objects", 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(".myvcs/commits", 0755); err != nil {
		return err
	}

	if err := os.WriteFile(".myvcs/HEAD", []byte(""), 0644); err != nil {
		return err
	}

	return nil
}

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
