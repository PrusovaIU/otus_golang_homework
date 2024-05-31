package dd_copy

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"

	"os"

	"github.com/schollz/progressbar/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

type File interface {
	Stat() (fs.FileInfo, error)
	Seek(int64, int) (int64, error)
}

func offsetPrepare(fromFile File, offset int64) error {
	fromFileStat, err := fromFile.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	if offset > fromFileStat.Size() {
		return ErrOffsetExceedsFileSize
	}
	_, err = fromFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}
	return nil
}

type BufferByteReader interface {
	ReadByte() (byte, error)
}

type BufferByteWriter interface {
	WriteByte(byte) error
	Flush() error
}

func readWrite(bufferReader BufferByteReader, bufferWriter BufferByteWriter, limit int64) error {
	bar := progressbar.Default(limit)
	for i := 0; true; i++ {
		if limit != 0 && int64(i) >= limit {
			break
		}
		ibyte, err := bufferReader.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		if err = bufferWriter.WriteByte(ibyte); err != nil {
			return err
		}
		if err := bar.Add(1); err != nil {
			fmt.Println("Cannot show bar")
		}
	}
	if err := bar.Finish(); err != nil {
		fmt.Println("Cannot show bar")
	}
	if err := bufferWriter.Flush(); err != nil {
		return err
	}
	return nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	if err = offsetPrepare(fromFile, offset); err != nil {
		fromFile.Close()
		return err
	}
	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	fromFileReader := bufio.NewReader(fromFile)
	toFileWriter := bufio.NewWriter(toFile)
	if err := readWrite(fromFileReader, toFileWriter, limit); err != nil {
		return err
	}
	if err := toFile.Close(); err != nil {
		return err
	}
	if err := fromFile.Close(); err != nil {
		return err
	}
	return nil
}
