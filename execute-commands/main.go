package main

import (
	"os"
	"os/exec"
)

func executeCommand(command string, argsArr []string) error {
	cmdObj := exec.Command(command, argsArr...)

	cmdObj.Stdout = os.Stdout
	cmdObj.Stderr = os.Stderr

	err := cmdObj.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	command := "ls"

	executeCommand(command, []string{"-l"})
}
