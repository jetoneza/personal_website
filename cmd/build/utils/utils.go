package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(name string, args ...string) {
	executeCommand("", name, args...)
}

func ExecuteCommandWithEnv(env string, name string, args ...string) {
	executeCommand(env, name, args...)
}

func HasCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)

	return err == nil
}

func executeCommand(env string, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if env != "" {
		newEnv := append(os.Environ(), env)
		cmd.Env = newEnv
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
