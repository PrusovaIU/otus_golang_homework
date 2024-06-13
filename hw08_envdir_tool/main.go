package main

import (
	"fmt"
	"os"
)

func main() {
	envDirPath := os.Args[1]
	command := os.Args[2]
	args := os.Args[3:]

	environment, err := ReadDir(envDirPath)
	if err == nil {
		RunCmd(command, args, environment)
		// fmt.Println("Exit code: ", exitCode)
	} else {
		fmt.Println(err)
	}

}
