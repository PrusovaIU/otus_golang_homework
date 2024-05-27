package copy

import (
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
	t.Run("test_success", func(t *testing.T) {
		stat_mock := &StatMock{}
		stat_mock.On("Size").Return(offset)
		file_mock := mocks.NewFile(t)
		file_mock.EXPECT().Stat().Return(stat_mock, nil)
		file_mock.EXPECT().Seek(int64(offset), io.SeekStart).Return(int64(0), nil)

		error := offsetPrepare(file_mock, int64(offset))
		require.Nil(t, error)

	})

}
