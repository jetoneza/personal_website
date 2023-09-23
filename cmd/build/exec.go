package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jetoneza/personal_website/cmd/build/utils"
	"github.com/jetoneza/personal_website/pkg/config"
)

const templatePath = "./web"

func buildProduction() {
	buildWeb()
	buildServer()
}

func runBackendProd() {
	log.Println("Running the go server...")
	utils.ExecuteCommand("./webapp")
}

func runFrontendProd() {
	log.Println("Running the svelete server...")

	env := fmt.Sprintf("PORT=%v", config.SVELTE_PORT)
	path := fmt.Sprintf("%v/build", templatePath)

	utils.ExecuteCommandWithEnv(env, "node", path)
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
	case "run:backend:prod":
		runBackendProd()
	case "run:frontend:prod":
		runFrontendProd()
	default:
		fmt.Printf("Invalid command '%v'\n", command)
	}
}
