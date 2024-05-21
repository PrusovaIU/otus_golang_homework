package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	// Place your code here.
}

func TestOpenFromFile(t *testing.T) {
	file_name := "test_file"
	file_content := "Package os provides a platform-independent interface to operating system functionality. The design is Unix-like, although the error handling is Go-like; failing calls return values of type error rather than error numbers. Often, more information is available within the error. For example, if a call that takes a file name fails, such as Open or Stat, the error will include the failing file name when printed and will be of type *PathError, which may be unpacked for more information."
	t.Run("success_opening", func(t *testing.T) {
		if _, err := os.Stat(file_name); err == nil {
			if err := os.Remove(file_name); err != nil {
				t.FailNow()
			}
		} else if !errors.Is(err, os.ErrNotExist) {
			t.FailNow()
		}
		file, err := os.Create(file_name)
		if err != nil {
			t.FailNow()
		}
		file.WriteString(file_content)
		file.Close()
		offset := int(len(file_content) / 2)
		_, err = openFromFile(file_name, int64(offset))
		require.Nil(t, err)
		os.Remove(file_name)
	})

	// t.Run("file_not_exist", func(t *testing.T) {
	// 	if _, err := os.Stat(); err == nil {

	// 	}
	// })
}
