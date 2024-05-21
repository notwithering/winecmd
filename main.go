package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var args = []string{"cmd.exe"}
	args = append(args, os.Args[1:]...)

	cmd := exec.Command("wine", args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = nil

	err := cmd.Run()

	if cmd.ProcessState.ExitCode() != 0 {
		fmt.Println(err)
	}
}
