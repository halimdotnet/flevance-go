package pcard

import "testing"

func TestGetIssuer(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  IssuerNetwork
	}{
		{"Visa", "4532015112345678", INVisa},
		{"Amex 34", "340000000000009", INAmex},
		{"Amex 37", "370000000000002", INAmex},
		{"JCB lower", "3528000000000000", INJCB},
		{"JCB upper", "3589000000000000", INJCB},
		{"MasterCard old range", "5100000000000000", INMastercard},
		{"MasterCard new range", "2221000000000000", INMastercard},
		{"Discover 65", "6500000000000000", INDiscover},
		{"Discover 6011", "6011000000000000", INDiscover},
		{"Discover 644-649", "6440000000000000", INDiscover},
		{"Discover 622126-622925", "6221260000000000", INDiscover},
		{"Unknown", "1234567890123456", ""},
		//{"Empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getIssuer(tt.input)
			if got != tt.want {
				t.Errorf("getIssuer() = %v, want %v", got, tt.want)
			}
		})
	}
}
