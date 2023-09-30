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
	log.Println("GOLANG: running production server")
	utils.ExecuteCommand("./bin/api")
}

func runFrontendProd() {
	log.Println("SVELTE: running production server")

	env := fmt.Sprintf("PORT=%v", config.SVELTE_PORT)
	path := fmt.Sprintf("%v/build", templatePath)

	utils.ExecuteCommandWithEnv(env, "node", path)
}

func buildServer() {
	log.Println("GOLANG: building production server")
	utils.ExecuteCommand("go", "build", "-o", "bin/api", "-v")
}

func runServer() {
	log.Println("GOLANG: running development server")
	utils.ExecuteCommand("go", "run", "main.go")
}

func buildWeb() {
	log.Println("SVELTE: building production project")
	installWebDependencies()

	if utils.HasCommand("pnpm") {
		utils.ExecuteCommand("pnpm", "run", "-C", templatePath, "build")
		return
	}

	utils.ExecuteCommand("npm", "run", "build", "--prefix ", templatePath)
}

func runWeb() {
	log.Println("SVELTE: running development project")
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
