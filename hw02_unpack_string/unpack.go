package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	var stringBuilder strings.Builder
	inputStringRunes := []rune(inputString)

	for currentIndex, currentRune := range inputStringRunes {
		isCurrentRuneDigit := unicode.IsDigit(currentRune)
		if currentIndex == 0 && isCurrentRuneDigit {
			return "", ErrInvalidString
		}

		nextRune, err := getRuneByIndex(inputStringRunes, currentIndex+1)
		if err != nil && isCurrentRuneDigit {
			continue
		}

		isNextRuneDigit := unicode.IsDigit(nextRune)
		if isCurrentRuneDigit && isNextRuneDigit {
			return "", ErrInvalidString
		}

		if isNextRuneDigit {
			repeatCount, err := strconv.Atoi(string(nextRune))
			if err != nil {
				return "", err
			}

			stringBuilder.WriteString(repeatStringByRune(currentRune, repeatCount))
		} else if !isCurrentRuneDigit {
			stringBuilder.WriteString(string(currentRune))
		}
	}

	return stringBuilder.String(), nil
}

func getRuneByIndex(inputStringRunes []rune, index int) (rune, error) {
	inputStringLength := len(inputStringRunes)
	if index >= inputStringLength {
		return 0, fmt.Errorf("индекс превышает размер строки")
	}

	return inputStringRunes[index], nil
}

func repeatStringByRune(r rune, repeatCount int) string {
	return strings.Repeat(string(r), repeatCount)
}
