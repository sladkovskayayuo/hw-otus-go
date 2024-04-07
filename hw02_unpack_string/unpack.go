package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {

	if str == "" {
		return "", nil
	}

	n := -1
	var result strings.Builder
	letter := rune(-1)

	for i, char := range str {
		if char > '9' || char < '0' {
			if letter == -1 {
				letter = char
			} else {
				result.WriteString(string(letter))
				letter = char
			}
		} else if char >= '0' && char <= '9' {
			if n == -1 && letter != -1 {
				n = (int(char) - '0')
			} else {
				return "", ErrInvalidString
			}
		} else {
			return "", ErrInvalidString
		}
		if n != -1 {
			result.WriteString(strings.Repeat(string(letter), n))
			n = -1
			letter = -1
		}

		if i == len(str)-1 && letter != -1 {
			result.WriteString(string(letter))
		}
	}

	return result.String(), nil
}
