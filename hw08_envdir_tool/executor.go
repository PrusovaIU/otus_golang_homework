package main

import (
	"fmt"
	"os"
	"os/exec"
)

func formEnv(dirEnv Environment) error {
	for name, value := range dirEnv {
		if value.NeedRemove {
			if err := os.Unsetenv(name); err != nil {
				return err
			}
		} else {
			if err := os.Setenv(name, value.Value); err != nil {
				return err
			}
		}
	}
	return nil
}

type Process interface {
	Start() error
	Wait() error
}

func processManage(process Process) error {
	if err := process.Start(); err != nil {
		return err
	}
	if err := process.Wait(); err != nil {
		return err
	}
	return nil
}

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd string, args []string, env Environment) (returnCode int) {
	checkErr := func(err error) int {
		fmt.Println(err)
		return 1
	}
	if err := formEnv(env); err != nil {
		return checkErr(err)
	}
	process := exec.Command(cmd, args...)
	process.Env = os.Environ()
	process.Stderr = os.Stderr
	process.Stdout = os.Stdout
	process.Stdin = os.Stdin
	if err := processManage(process); err != nil {
		return checkErr(err)
	}
	return process.ProcessState.ExitCode()
}
