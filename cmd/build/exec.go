package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jetoneza/personal_website/cmd/build/utils"
)

const templatePath = "./web"

func buildProduction() {
	buildWeb()
	buildServer()
}

func buildServer() {
	log.Println("Building the go project...")
	utils.ExecuteCommand("go", "build", "-o", "webapp", "-v")
}

func runServer() {
	log.Println("Running the go project...")
	utils.ExecuteCommand("go", "run", "main.go")
}

func buildWeb() {
	log.Println("Building sveltekit...")
	installWebDependencies()

	if utils.HasCommand("pnpm") {
		utils.ExecuteCommand("pnpm", "run", "-C", templatePath, "build")
		return
	}

	utils.ExecuteCommand("npm", "run", "build", "--prefix ", templatePath)
}

func runWeb() {
	log.Println("Running sveltekit...")
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
	case "build:prod":
		buildProduction()
	default:
		fmt.Printf("Invalid command '%v'\n", command)
	}
}
