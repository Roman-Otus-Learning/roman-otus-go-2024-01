package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	sourceFile, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer sourceFile.Close() //nolint:golint

	sourceFileSize, err := getFileSize(sourceFile)
	if err != nil {
		return err
	}

	err = makeOffset(offset, sourceFileSize, sourceFile)
	if err != nil {
		return err
	}

	if limit == 0 || limit > (sourceFileSize-offset) {
		limit = sourceFileSize - offset
	}

	progressBar := createProgressBar(limit)
	progressBar.Start()
	defer progressBar.Finish()
	sourceReader := progressBar.NewProxyReader(sourceFile)

	destinationFile, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %w", err)
	}
	defer destinationFile.Close() //nolint:golint

	err = copyFile(sourceReader, destinationFile, limit)
	if err != nil {
		return err
	}

	return nil
}

func copyFile(sourceReader *pb.Reader, destinationFile *os.File, limit int64) error {
	_, err := io.CopyN(destinationFile, sourceReader, limit)

	if err != nil && !errors.Is(err, io.EOF) {
		removeErr := os.Remove(destinationFile.Name())
		if removeErr != nil {
			return fmt.Errorf("не удалось удалить файл в результате ошибки копирования: %w", removeErr)
		}

		return fmt.Errorf("не удалось копировать файл: %w", err)
	}

	return nil
}

func createProgressBar(total int64) *pb.ProgressBar {
	bar := pb.New(int(total)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.ShowSpeed = true

	return bar
}

func getFileSize(file *os.File) (size int64, err error) {
	sourceFileStats, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf("не удалось получить информацию о файле: %w", err)
	}

	fileSize := sourceFileStats.Size()
	if fileSize <= 0 {
		return 0, ErrUnsupportedFile
	}

	return fileSize, nil
}

func makeOffset(offset, sourceFileSize int64, file *os.File) error {
	if offset == 0 {
		return nil
	}

	if offset > sourceFileSize {
		return ErrOffsetExceedsFileSize
	}

	_, err := file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("не удалось применить offset к файлу: %w", err)
	}

	return nil
}
