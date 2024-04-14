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

	for _, char := range str {

		if char > '9' || char < '0' {

			if letter == -1 {

				letter = char

			} else {

				result.WriteString(string(letter))

				letter = char

			}

		} else {

			if n == -1 && letter != -1 {

				n = (int(char) - '0')

			} else {

				return "", ErrInvalidString

			}

		}

		if n != -1 {

			result.WriteString(strings.Repeat(string(letter), n))

			n = -1

			letter = -1

		}

	}

	if letter != -1 {

		result.WriteString(string(letter))

	}

	return result.String(), nil

}
