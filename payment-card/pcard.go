package pcard

import (
	"strconv"
	"strings"
	"time"
)

func ValidateNumber(input string) bool {
	n := len(input)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[n-1-i] = int(input[i] - '0')
	}

	var sum int
	for i, v := range numbers {

		t := v

		if i > 0 && i%2 == 0 {
			t = v * 2
			if t > 9 {
				t -= 9
			}
		}

		sum += t
	}

	return sum%10 == 0
}

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

func ValidateExpiry(input string) (bool, time.Time, error) {
	var exp bool
	var expd time.Time
	var err error

	in := ReplaceChar(input)
	if len(in) != 4 && len(in) != 6 {
		return exp, expd, ErrInvalidExpiryFormat
	}

	var mStr, yStr = in[0:2], in[2:]
	var m, y int

	if m, err = strconv.Atoi(mStr); err != nil {
		return exp, expd, ErrInvalidExpiryFormat
	}

	if y, err = strconv.Atoi(yStr); err != nil {
		return exp, expd, ErrInvalidExpiryFormat
	}

	if len(yStr) == 2 {
		y = (time.Now().Year()/100)*100 + y
	}

	expd = time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC).
		AddDate(0, 1, 0).
		Add(-time.Second)

	dCur := time.Now()
	if expd.Before(dCur) {
		exp = true
	}

	return exp, expd, nil
}
