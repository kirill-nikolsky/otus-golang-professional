package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const BackslashRune = 92

var (
	ErrInvalidString      = errors.New("invalid string")
	ErrPositionOutOfRange = errors.New("position is out of range")
)

func ConvertStringToRunesSlice(tc string) []rune {
	return []rune(tc)
}

func GetRuneAtPosition(runes []rune, position int) (rune, error) {
	if position < 0 || position >= len(runes) {
		return 0, ErrPositionOutOfRange
	}

	return runes[position], nil
}

func Unpack(tc string) (string, error) {
	runes := ConvertStringToRunesSlice(tc)
	runesLen := utf8.RuneCountInString(tc)
	backslashMarks := make(map[int]bool, runesLen)

	var result strings.Builder

	for i, backslashCounter := 0, 0; i < runesLen; i++ {
		currentRune, _ := GetRuneAtPosition(runes, i)

		if unicode.IsDigit(currentRune) {
			// String starts with digit => error case
			if i == 0 {
				return "", ErrInvalidString
			}

			prevRune, _ := GetRuneAtPosition(runes, i-1)

			// Two digits in a row => error case
			if unicode.IsDigit(prevRune) && !backslashMarks[i-1] {
				return "", ErrInvalidString
			}

			if prevRune != BackslashRune {
				continue
			}
		}

		if currentRune == BackslashRune {
			backslashCounter++
			continue
		}

		if backslashCounter != 0 {
			backslashMarks[i] = true
			backslashCounter = 0
		}

		// by default assume next rune = '1'
		nextRune := '1'

		// check for slice range, and get next rune
		if i < runesLen-1 {
			nextRune, _ = GetRuneAtPosition(runes, i+1)
		}

		// by default repeat currentRune once
		var repeatCounter int64 = 1

		if unicode.IsDigit(nextRune) {
			repeatCounter, _ = strconv.ParseInt(string(nextRune), 10, 64)

			// if repeatCounter == 0, ignore currentRune
			if repeatCounter == 0 {
				continue
			}
		}

		if repeatCounter > 0 {
			result.WriteString(strings.Repeat(string(currentRune), int(repeatCounter)))
		}
	}

	return result.String(), nil
}
