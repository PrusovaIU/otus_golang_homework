package copy

import (
	"bufio"
	"errors"
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
	for i := 0; i < int(limit); i++ {
		ibyte, err := bufferReader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		err = bufferWriter.WriteByte(ibyte)
		if err != nil {
			return err
		}
		bar.Add(1)
	}
	bar.Finish()
	err := bufferWriter.Flush()
	if err != nil {
		return err
	}
	return nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	err = offsetPrepare(fromFile, offset)
	if err != nil {
		fromFile.Close()
		return err
	}
	toFile, err := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	fromFileReader := bufio.NewReader(fromFile)
	toFileWriter := bufio.NewWriter(toFile)
	readWrite(fromFileReader, toFileWriter, limit)
	toFile.Close()
	fromFile.Close()
	return nil
}
