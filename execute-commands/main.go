package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	command := "date"

	cmdObj := exec.Command(command)

	cmdObj.Stdout = os.Stdout
	cmdObj.Stderr = os.Stderr

	err := cmdObj.Run()
	if err != nil {
		log.Fatal(err)
	}
}
