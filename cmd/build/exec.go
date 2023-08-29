package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Colors
var green = color.New(color.FgGreen)

func build() {
	green.Println("Building the go project...")
	ExecuteCommand("go", "build", "-o", "webapp", "-v")
}

func run() {
	green.Println("Running the go project...")
	ExecuteCommand("go", "run", "main.go")
}

// TODO: Add svelte-kit build/run commands

func main() {
	command := os.Args[1]

	switch command {
	case "build":
		build()
	case "run":
		run()
	default:
		fmt.Printf("Invalid command '%v'\n", command)
	}
}
