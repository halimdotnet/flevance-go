package main

import (
	"fmt"

	pcard "go.halimdotnet.dev/flevance-go/payment-card"
)

var ()

func main() {
	//CardNumberValidation()
	PaymentCardValidation()
}

func PaymentCardValidation() {
	d, err := pcard.ValidatePaymentCard("4532 0151 1234 5678", "10/29")
	if err != nil {
		panic(err)
	}

	fmt.Println("Valid Card Number: ", d.ValidNumber)
	fmt.Println("Valid Expiry Date: ", d.ValidExpiry)
	fmt.Println("Valid Issuer Network: ", d.ValidIssuer)
	fmt.Println("Card Number ", d.Data.Number)
	fmt.Println("Issuer Name: ", d.Data.Issuer)
	if d.ValidExpiry == true {
		fmt.Println("Valid until: ", d.Data.Expiry)
	}

}

func CardNumberValidation(n string) {

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
