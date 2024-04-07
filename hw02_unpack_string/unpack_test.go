package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {

	tests := []struct {
		input string

		expected string
	}{

		{input: "a4bc2d5e", expected: "aaaabccddddde"},

		{input: "abccd", expected: "abccd"},

		{input: "", expected: ""},

		{input: "aaa0b", expected: "aab"},

		// uncomment if task with asterisk completed

		// {input: `qwe\4\5`, expected: `qwe45`},

		// {input: `qwe\45`, expected: `qwe44444`},

		// {input: `qwe\\5`, expected: `qwe\\\\\`},

		// {input: `qwe\\\3`, expected: `qwe\3`},

	}

	TestsGroupNoErrorRuner(tests, t)

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

func TestUnpackStringWithSpecialChar(t *testing.T) {

	tests := []struct {
		input string

		expected string
	}{

		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},

		{input: "d\t4", expected: "d\t\t\t\t"},

		{input: "d\n", expected: "d\n"},
	}

	TestsGroupNoErrorRuner(tests, t)

}

func TestsGroupNoErrorRuner(tests []struct {
	input string

	expected string
}, t *testing.T) {

	for _, tc := range tests {

		tc := tc

		t.Run(tc.input, func(t *testing.T) {

			result, err := Unpack(tc.input)

			require.NoError(t, err)

			require.Equal(t, tc.expected, result)

		})

	}

}
