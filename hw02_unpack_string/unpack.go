package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var prev string
	var skipNext bool

	for _, r := range s {
		if string(r) == "\\" && !skipNext {
			skipNext = true
			continue
		}

		if unicode.IsDigit(r) && !skipNext {
			if prev == "" {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(string(r))
			if count == 0 {
				prev = ""
				continue
			}

			_, err := result.WriteString(strings.Repeat(prev, count))
			if err != nil {
				return "", ErrInvalidString
			}

			prev = ""
			continue
		}

		if prev != "" {
			result.WriteString(prev)
		}

		skipNext = false
		prev = string(r)
	}

	if prev != "" {
		result.WriteString(prev)
	}

	return result.String(), nil
}
