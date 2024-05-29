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
