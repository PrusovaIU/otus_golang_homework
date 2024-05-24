package main

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func openFromFile(path string, offset int64) (*os.File, error) {
	fromFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fromFileStat, err := fromFile.Stat()
	if err != nil {
		return nil, ErrUnsupportedFile
	}
	if offset > fromFileStat.Size() {
		fromFile.Close()
		return nil, ErrOffsetExceedsFileSize
	}
	_, err = fromFile.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return fromFile, nil
}

func copy(fromFile, toFile *os.File, limit int64) error {
	fromFileReader := bufio.NewReader(fromFile)
	toFileWriter := bufio.NewWriter(toFile)
	bar := progressbar.Default(limit)
	for i := 0; i < int(limit); i++ {
		ibyte, err := fromFileReader.ReadByte()
		if err != nil {
			return err
		}
		err = toFileWriter.WriteByte(ibyte)
		if err != nil {
			return err
		}
		bar.Add(1)
	}
	toFileWriter.Flush()
	return nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := openFromFile(fromPath, offset)
	if err != nil {
		return err
	}
	toFile, err := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	copy(fromFile, toFile, limit)
	fromFile.Close()
	toFile.Close()
	return nil
}
