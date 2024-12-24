package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "🙃0", expected: ""},
		{input: "aaф0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4`, expected: `qwe4`},
		{input: `\45`, expected: `44444`},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestConvertStringToRunesSlice(t *testing.T) {
	tests := []struct {
		input    string
		expected []rune
	}{
		{input: "a4bc2d5e", expected: []rune("a4bc2d5e")},
		{input: "abccd", expected: []rune("abccd")},
		{input: "", expected: []rune("")},
		{input: "aaa0b", expected: []rune("aaa0b")},
		{input: "🙃0", expected: []rune("🙃0")},
		{input: "aaф0b", expected: []rune("aaф0b")},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := ConvertStringToRunesSlice(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestGetRuneAtPosition(t *testing.T) {
	tests := []struct {
		input    []rune
		position int
		expected rune
	}{
		{input: []rune("a4bc2d5e"), position: 0, expected: 'a'},
		{input: []rune("abccd"), position: 2, expected: 'c'},
		{input: []rune("aaa0b"), position: 3, expected: '0'},
		{input: []rune("🙃0"), position: 0, expected: '🙃'},
		{input: []rune("aaф0b"), position: 2, expected: 'ф'},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(string(tc.input), func(t *testing.T) {
			result, err := GetRuneAtPosition(tc.input, tc.position)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}

	t.Run("position is out of range upper", func(t *testing.T) {
		_, err := GetRuneAtPosition([]rune("a4bc2d5e"), 10)
		require.ErrorIs(t, err, ErrPositionOutOfRange)
	})

	t.Run("position is out of range negative", func(t *testing.T) {
		_, err := GetRuneAtPosition([]rune("a4bc2d5e"), -1)
		require.ErrorIs(t, err, ErrPositionOutOfRange)
	})

	t.Run("position zero in zero-length", func(t *testing.T) {
		_, err := GetRuneAtPosition([]rune(""), 0)
		require.ErrorIs(t, err, ErrPositionOutOfRange)
	})
}
