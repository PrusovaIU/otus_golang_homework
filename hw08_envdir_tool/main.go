package main

import (
	"fmt"
	"os"
)

func main() {
	envDirPath := os.Args[1]
	command := os.Args[2]
	args := os.Args[3:]

	fmt.Println(envDirPath, command, args)
	fmt.Println(ReadDir(envDirPath))
}
