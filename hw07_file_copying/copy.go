package main

import (
	"errors"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func openFromFile(path string, offset int64) (*os.File, error) {
	fromFile, err := os.Open(path)
	if err != nil {
		return nil, ErrUnsupportedFile
	}
	fromFileStat, err := fromFile.Stat()
	if err != nil {
		return nil, ErrUnsupportedFile
	}
	if offset > fromFileStat.Size() {
		fromFile.Close()
		return nil, ErrOffsetExceedsFileSize
	}
	return fromFile, nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	// fromFile, err := openFromFile(fromPath, offset)
	// if err != nil {
	// 	return err
	// }

	return nil
}
