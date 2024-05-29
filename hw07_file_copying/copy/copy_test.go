package copy

import (
	"errors"
	"io"
	"io/fs"
	"testing"

	"otus_golang_homework/hw07_file_copying/copy/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	// Place your code here.
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
