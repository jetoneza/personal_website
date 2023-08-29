package main

import (
	"fmt"
	"os"

  "github.com/fatih/color"
)

// Colors
var green = color.New(color.FgGreen)

func Build() {
  green.Println("Building the project...")
  ExecuteCommand("go", "build", "-o", "webapp", "-v")
}

// TODO: Add svelte-kit build/run commands

func main() {
	command := os.Args[1]

	switch command {
	case "build":
		Build()
	default:
		fmt.Printf("Invalid command '%v'\n", command)
	}
}
