package main

import (
	"fmt"
	"os"
)

func initialise(name string) error {
	fmt.Println("initialising project " + name)

	err := os.MkdirAll(".myvcs", 0755)
	if err != nil {
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

	fmt.Println("initialised")

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Init requires a project name")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		name := os.Args[2]
		initialise(name)
	}
}
