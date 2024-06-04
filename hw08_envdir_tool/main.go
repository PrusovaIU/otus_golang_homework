package main

import (
	"fmt"
	"os"
)

func main() {
	env_dir_path := os.Args[1]
	command := os.Args[2]
	args := os.Args[3:]

	fmt.Println(env_dir_path, command, args)
	fmt.Println(ReadDir(env_dir_path))
}
