package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jetoneza/personal_website/cmd/build/utils"
)

// Colors
var (
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
)

const templatePath = "./internal/template"

func buildServer() {
	green.Println("Building the go project...")
	utils.ExecuteCommand("go", "build", "-o", "webapp", "-v")
}

func runServer() {
	green.Println("Running the go project...")
	utils.ExecuteCommand("go", "run", "main.go")
}

func buildWeb() {
	green.Println("Building sveltekit...")
	installWebDependencies()

	if utils.HasCommand("pnpm") {
		utils.ExecuteCommand("pnpm", "run", "-C", templatePath, "build")
		return
	}

	utils.ExecuteCommand("npm", "run", "build", "--prefix ", templatePath)
}

func runWeb() {
	green.Println("Running sveltekit...")
	installWebDependencies()

	if utils.HasCommand("pnpm") {
		utils.ExecuteCommand("pnpm", "run", "-C", templatePath, "dev")
		return
	}

	utils.ExecuteCommand("npm", "run", "dev", "--prefix ", templatePath)
}

func installWebDependencies() {
	if utils.HasCommand("pnpm") {
		utils.ExecuteCommand("pnpm", "install", "-C", templatePath)
		return
	}

	utils.ExecuteCommand("npm", "install", "--prefix", templatePath)
}

// TODO: Add svelte-kit build/run commands

func main() {
	command := os.Args[1]

	switch command {
	case "build:web":
		buildWeb()
	case "run:web":
		runWeb()
	case "build":
		buildServer()
	case "run":
		runServer()
	default:
		fmt.Printf("Invalid command '%v'\n", command)
	}
}
