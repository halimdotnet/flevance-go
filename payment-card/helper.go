package pcard

import "strings"

func ReplaceChar(input string) string {
	var sb strings.Builder
	sb.Grow(len(input))
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c >= '0' && c <= '9' {
			sb.WriteByte(c)
		}
	}

	return sb.String()
}

func Digits(input string) error {
	l := len(input)
	if l >= MIN_DIGITS && l <= MAX_DIGITS {
		return nil
	}

	return ErrInvalidTotalDigits
}
