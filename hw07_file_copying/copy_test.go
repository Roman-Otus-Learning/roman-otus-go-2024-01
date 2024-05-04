package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("Успешное копирование", func(t *testing.T) {
		err := Copy("testdata/input_short.txt", "out.txt", 0, 0)

		require.NoError(t, err)
		file, err := os.Open("out.txt")
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		require.NoError(t, err)
		err = os.Remove("out.txt")
		require.NoError(t, err)
	})

	t.Run("Несуществующий файл", func(t *testing.T) {
		err := Copy("not-existed-file.txt", "out.txt", 1000, 0)
		require.Error(t, err)
	})

	t.Run("Ошибка превышения офсета", func(t *testing.T) {
		err := Copy("testdata/input_short.txt", "out.txt", 1000, 0)
		require.Truef(t, errors.Is(err, ErrOffsetExceedsFileSize), "actual error %q", err)
	})

	t.Run("Ошибка копирования из-за пустого файла", func(t *testing.T) {
		err := Copy("/dev/urandom", "out.txt", 1000, 0)
		require.Truef(t, errors.Is(err, ErrUnsupportedFile), "actual error %q", err)
	})
}
