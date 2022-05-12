package main

import (
	"os"
	"os/exec"
)

func executeCommand(command string, argsArr []string) error {
	cmdObj := exec.Command(command, argsArr...)

	cmdObj.Stdout = os.Stdout
	cmdObj.Stderr = os.Stderr
	cmdObj.Stdin = os.Stdin

	err := cmdObj.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// ifconfig wlp4s0 down
	executeCommand("sudo", []string{"ifconfig", "wlp4s0", "down"})
	// ifconfig wlp4s0 hw ether 00:11:22:33:44:55
	executeCommand("sudo", []string{"ifconfig", "wlp4s0", "hw", "ether", "00:11:22:33:44:55"})
	// ifconfig wlp4s0 up
	executeCommand("sudo", []string{"ifconfig", "wlp4s0", "up"})
}
