package copy

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"os"
	"testing"

	"otus_golang_homework/hw07_file_copying/copy/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// func TestCopy(t *testing.T) {

// }

type CopyTestSuite struct {
	suite.Suite
	fromFilePath, toFilePath string
	fileContent              string
	toFile                   *os.File
}

func (suite *CopyTestSuite) deleteFile(filePath string) {
	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			panic(err)
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}
}

func (suite *CopyTestSuite) SetupTest() {
	suite.fileContent = "Package suite contains logic for creating testing suite structs and running the methods on those structs as tests. The most useful piece of this package is that you can create setup/teardown methods on your testing suites, which will run before/after the whole suite or individual tests (depending on which interface(s) you implement)."
	suite.fromFilePath = "from_file_test"
	suite.toFilePath = "to_file_test"
	fromFile, err := os.OpenFile(suite.fromFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	if _, err := fromFile.WriteString(suite.fileContent); err != nil {
		panic(err)
	}
	fromFile.Close()
	suite.deleteFile(suite.toFilePath)
}

func (suite *CopyTestSuite) TearDownTest() {
	suite.toFile.Close()
	suite.deleteFile(suite.fromFilePath)
	suite.deleteFile(suite.toFilePath)
}

func (suite *CopyTestSuite) TestSuccess() {
	limit := 10
	err := Copy(suite.fromFilePath, suite.toFilePath, 0, int64(limit))
	require.NoError(suite.T(), err)
	toFile, err := os.Open(suite.toFilePath)
	suite.toFile = toFile
	require.NoError(suite.T(), err)
	buffer := bufio.NewReader(toFile)
	fileContent := []byte{}
	for i := 0; i < limit; i++ {
		ibyte, err := buffer.ReadByte()
		require.NoError(suite.T(), err)
		fileContent = append(fileContent, ibyte)
	}
	fileData := string(fileContent[:])
	require.Equal(suite.T(), suite.fileContent[:10], fileData)
}

func TestCopy(t *testing.T) {
	suite.Run(t, new(CopyTestSuite))
}

type StatMock struct {
	mock.Mock
	fs.FileInfo
}

func (o *StatMock) Size() int64 {
	args := o.Called()
	return int64(args.Int(0))
}

func TestOffsetPrepare(t *testing.T) {
	offset := 100

	cases := []struct {
		name string
		size int
	}{
		{name: "size_eq_offset", size: offset},
		{name: "size_gt_offset", size: offset + 1},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			stat_mock := &StatMock{}
			stat_mock.On("Size").Return(tc.size)
			file_mock := mocks.NewFile(t)
			file_mock.EXPECT().Stat().Return(stat_mock, nil)
			file_mock.EXPECT().Seek(int64(offset), io.SeekStart).Return(int64(0), nil)

			err := offsetPrepare(file_mock, int64(offset))
			require.Nil(t, err)
		})
	}

	t.Run("size_lt_offset", func(t *testing.T) {
		stat_mock := &StatMock{}
		stat_mock.On("Size").Return(offset - 1)
		file_mock := mocks.NewFile(t)
		file_mock.EXPECT().Stat().Return(stat_mock, nil)

		err := offsetPrepare(file_mock, int64(offset))
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("stat_error", func(t *testing.T) {
		file_mock := mocks.NewFile(t)
		file_mock.EXPECT().Stat().Return(nil, errors.New("Test error"))

		err := offsetPrepare(file_mock, int64(offset))
		require.ErrorIs(t, err, ErrUnsupportedFile)
	})

	t.Run("seek_error", func(t *testing.T) {
		stat_mock := &StatMock{}
		stat_mock.On("Size").Return(offset)
		file_mock := mocks.NewFile(t)
		file_mock.EXPECT().Stat().Return(stat_mock, nil)
		returned_err := errors.New("Test error")
		file_mock.EXPECT().Seek(int64(offset), io.SeekStart).Return(0, returned_err)

		err := offsetPrepare(file_mock, int64(offset))
		require.ErrorIs(t, err, returned_err)
	})
}

func TestReadWrite(t *testing.T) {
	offset := 10
	read_byte := byte(1)
	returned_err := errors.New("Test error")

	t.Run("stop_by_limit", func(t *testing.T) {
		buffer_reader := mocks.NewBufferByteReader(t)
		buffer_reader.EXPECT().ReadByte().Return(read_byte, nil)
		buffer_writer := mocks.NewBufferByteWriter(t)
		buffer_writer.EXPECT().WriteByte(read_byte).Return(nil)
		buffer_writer.EXPECT().Flush().Return(nil)

		err := readWrite(buffer_reader, buffer_writer, int64(offset))
		require.NoError(t, err)

		buffer_reader.AssertNumberOfCalls(t, "ReadByte", offset)
		buffer_writer.AssertNumberOfCalls(t, "WriteByte", offset)
	})

	t.Run("stop_by_eof", func(t *testing.T) {
		bytes_in_file := int(offset / 2)
		buffer_reader := mocks.NewBufferByteReader(t)
		for i := 0; i < bytes_in_file; i++ {
			buffer_reader.EXPECT().ReadByte().Return(read_byte, nil).Once()
		}
		buffer_reader.EXPECT().ReadByte().Return(0, io.EOF).Once()
		buffer_writer := mocks.NewBufferByteWriter(t)
		buffer_writer.EXPECT().WriteByte(read_byte).Return(nil)
		buffer_writer.EXPECT().Flush().Return(nil)

		err := readWrite(buffer_reader, buffer_writer, int64(offset))
		require.NoError(t, err)

		buffer_reader.AssertNumberOfCalls(t, "ReadByte", bytes_in_file+1)
		buffer_writer.AssertNumberOfCalls(t, "WriteByte", bytes_in_file)
	})

	t.Run("read_err", func(t *testing.T) {
		buffer_reader := mocks.NewBufferByteReader(t)
		buffer_reader.EXPECT().ReadByte().Return(0, returned_err)
		buffer_writer := mocks.NewBufferByteWriter(t)

		err := readWrite(buffer_reader, buffer_writer, int64(offset))
		require.ErrorIs(t, err, returned_err)

		buffer_writer.AssertNotCalled(t, "WriteByte")
		buffer_writer.AssertNotCalled(t, "Flush")
	})

	t.Run("write_err", func(t *testing.T) {
		buffer_reader := mocks.NewBufferByteReader(t)
		buffer_reader.EXPECT().ReadByte().Return(read_byte, nil)
		buffer_writer := mocks.NewBufferByteWriter(t)
		buffer_writer.EXPECT().WriteByte(read_byte).Return(returned_err)

		err := readWrite(buffer_reader, buffer_writer, int64(offset))
		require.ErrorIs(t, err, returned_err)

		buffer_writer.AssertNotCalled(t, "Flush")
	})

	t.Run("flush_err", func(t *testing.T) {
		buffer_reader := mocks.NewBufferByteReader(t)
		buffer_reader.EXPECT().ReadByte().Return(read_byte, nil)
		buffer_writer := mocks.NewBufferByteWriter(t)
		buffer_writer.EXPECT().WriteByte(read_byte).Return(nil)
		buffer_writer.EXPECT().Flush().Return(returned_err)

		err := readWrite(buffer_reader, buffer_writer, int64(offset))
		require.ErrorIs(t, err, returned_err)
	})
}
