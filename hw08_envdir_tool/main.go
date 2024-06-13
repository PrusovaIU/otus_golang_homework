package main

import (
	"fmt"
	"os"
)

func main() {
	envDirPath := os.Args[1]
	command := os.Args[2]
	args := os.Args[3:]

	var exitCode int
	environment, err := ReadDir(envDirPath)
	if err == nil {
		exitCode = RunCmd(command, args, environment)
	} else {
		fmt.Println(err)
	}
	os.Exit(exitCode)
}
