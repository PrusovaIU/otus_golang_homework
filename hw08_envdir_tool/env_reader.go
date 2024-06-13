package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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

// Read the file and form value for env.
func readFile(file EnvFile) (EnvValue, error) {
	line, _, err := file.ReadLine()
	if err != nil && !errors.Is(err, io.EOF) {
		return NewEnvValue("", false), err
	}
	forbiddenSymbols := regexp.MustCompile(`(\s+$)`)
	value := forbiddenSymbols.ReplaceAllString(string(line), "")
	value = strings.ReplaceAll(value, "\x00", "\n")
	if len(value) == 0 {
		return NewEnvValue("", true), nil
	}
	return NewEnvValue(value, false), nil
}

type EnvFileInfo interface {
	Name() string
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
			envName := info.Name()
			if strings.Contains(envName, "=") {
				fmt.Println("Name of file can not contain \"=\"")
				return errors.New("file name error")
			}
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			envValue, err := readFile(bufio.NewReader(file))
			if err != nil {
				return err
			}
			envMap[envName] = envValue
		}
		return nil
	})
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}
	return envMap, nil
}
