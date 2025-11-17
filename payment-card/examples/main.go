package main

import (
	"fmt"

	pcard "go.halimdotnet.dev/flevance-go/payment-card"
)

func main() {
	CardNumberValidation()
	CardExpiryValidation()
}

func CardExpiryValidation() {
	var dates = []string{"12-20", "09/2021", "01-30"}

	for _, date := range dates {
		v, d, _ := pcard.ValidateExpiry(date)
		fmt.Printf("%s: [expired:%+v; %+v]\n", date, v, d)
	}

}

func CardNumberValidation() {
	var numbers = []string{"4532 0151 1234 5678"}

	for _, number := range numbers {
		num := pcard.ReplaceChar(number)

		err := pcard.Digits(num)
		if err != nil {
			fmt.Printf("%s Error: %v\n", number, err)
			continue
		}

		if pcard.ValidateNumber(num) {
			fmt.Printf("%s is valid\n", number)
		} else {
			fmt.Printf("%s is invalid\n", number)
		}

	}
}
