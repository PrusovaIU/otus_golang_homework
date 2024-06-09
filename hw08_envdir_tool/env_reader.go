package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func NewEnvValue(value string, needRemove bool) EnvValue {
	return EnvValue{
		Value:      value,
		NeedRemove: needRemove,
	}
}

type EnvFile interface {
	ReadLine() (line []byte, isPrefix bool, err error)
}

// Read the file and form value for env
func readFile(file EnvFile) (EnvValue, error) {
	line, _, err := file.ReadLine()
	if err != nil {
		return NewEnvValue("", false), err
	}
	if len(line) == 0 {
		return NewEnvValue("", true), nil
	} else {
		return NewEnvValue(string(line), false), nil
	}
}

type EnvFileInfo interface {
	Name() string
}

func getParametername(info EnvFileInfo) string {
	fileName := info.Name()
	forbidden_symbols := regexp.MustCompile(`[= ]`)
	fileName = forbidden_symbols.ReplaceAllString(fileName, "")
	return strings.ReplaceAll(fileName, "0x00", "")
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	envMap := Environment{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if path != dir {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			envValue, err := readFile(bufio.NewReader(file))
			if err != nil {
				return err
			}
			envName := getParametername(info)
			envMap[envName] = envValue
		}
		return nil
	})
	return nil, err
}
