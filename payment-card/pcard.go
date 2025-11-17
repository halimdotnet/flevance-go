package pcard

import (
	"strconv"
	"time"
)

type PaymentCard struct {
	ValidNumber bool
	ValidIssuer bool
	ValidExpiry bool
	Data        struct {
		Number string
		Issuer string
		Expiry time.Time
	}
}

func ValidatePaymentCard(num string, exp string) (result PaymentCard, err error) {
	n := ReplaceChar(num)
	if len(n) < 1 {
		return result, ErrInvalidTotalDigits
	}

	err = Digits(n)
	if err != nil {
		return result, err
	}

	var vnRes, veRes, viRes bool
	var riRes string
	var reRes time.Time

	vnRes = ValidateNumber(n)
	if vnRes == false {
		return result, ErrInvalidNumber
	}

	viRes, riRes, err = validateIssuerNetwork(n)
	if err != nil {
		return result, err
	}

	veRes, reRes, err = validateExpiry(exp)
	if err != nil {
		return result, err
	}

	result.ValidNumber = vnRes
	result.ValidIssuer = viRes
	result.ValidExpiry = veRes
	result.Data.Number = n
	result.Data.Issuer = riRes
	result.Data.Expiry = reRes

	return result, nil

}

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

func validateIssuerNetwork(input string) (bool, string, error) {

	issuer := getIssuer(input)
	if issuer == "" {
		return false, "", ErrIssuerNotFound
	}

	return true, string(issuer), nil

}

func validateExpiry(input string) (bool, time.Time, error) {
	var exp bool
	var expd time.Time
	var err error

	in := ReplaceChar(input)
	if len(in) != 4 && len(in) != 6 {
		return exp, expd, ErrInvalidExpiryDate
	}

	var mStr, yStr = in[0:2], in[2:]
	var m, y int

	if m, err = strconv.Atoi(mStr); err != nil {
		return exp, expd, ErrInvalidExpiryDate
	}

	if y, err = strconv.Atoi(yStr); err != nil {
		return exp, expd, ErrInvalidExpiryDate
	}

	if len(yStr) == 2 {
		y = (time.Now().Year()/100)*100 + y
	}

	expd = time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC).
		AddDate(0, 1, 0).
		Add(-time.Second)

	dCur := time.Now()
	if expd.After(dCur) {
		exp = true
	}

	return exp, expd, nil
}
