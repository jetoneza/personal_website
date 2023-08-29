package main

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(name string, args ...string) {
  cmd := exec.Command(name, args...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err := cmd.Run()

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
