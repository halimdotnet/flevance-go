package pcard

import (
	"errors"
	"testing"
)

func TestValidateIssuerNetwork(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantValid  bool
		wantIssuer string
		wantErr    error
	}{
		// Visa Tests
		{
			name:       "Visa - valid prefix 4",
			input:      "4532015112345678",
			wantValid:  true,
			wantIssuer: "Visa",
			wantErr:    nil,
		},
		{
			name:       "Visa - valid prefix 4 (short)",
			input:      "4111111111111",
			wantValid:  true,
			wantIssuer: "Visa",
			wantErr:    nil,
		},
		{
			name:       "Visa - single digit 4",
			input:      "4",
			wantValid:  true,
			wantIssuer: "Visa",
			wantErr:    nil,
		},

		// American Express Tests
		{
			name:       "Amex - valid prefix 34",
			input:      "340000000000009",
			wantValid:  true,
			wantIssuer: "American Express",
			wantErr:    nil,
		},
		{
			name:       "Amex - valid prefix 37",
			input:      "370000000000002",
			wantValid:  true,
			wantIssuer: "American Express",
			wantErr:    nil,
		},
		{
			name:       "Amex - prefix 34 only",
			input:      "34",
			wantValid:  true,
			wantIssuer: "American Express",
			wantErr:    nil,
		},
		{
			name:       "Amex - prefix 37 only",
			input:      "37",
			wantValid:  true,
			wantIssuer: "American Express",
			wantErr:    nil,
		},

		// JCB Tests
		{
			name:       "JCB - valid prefix 3528 (lower bound)",
			input:      "3528000000000000",
			wantValid:  true,
			wantIssuer: "JCB",
			wantErr:    nil,
		},
		{
			name:       "JCB - valid prefix 3589 (upper bound)",
			input:      "3589000000000000",
			wantValid:  true,
			wantIssuer: "JCB",
			wantErr:    nil,
		},
		{
			name:       "JCB - valid prefix 3550 (middle)",
			input:      "3550000000000000",
			wantValid:  true,
			wantIssuer: "JCB",
			wantErr:    nil,
		},
		{
			name:       "JCB - invalid prefix 3527 (below range)",
			input:      "3527000000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
		{
			name:       "JCB - invalid prefix 3590 (above range)",
			input:      "3590000000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},

		// MasterCard Tests
		{
			name:       "MasterCard - valid prefix 51 (lower bound old range)",
			input:      "5100000000000000",
			wantValid:  true,
			wantIssuer: "Mastercard",
			wantErr:    nil,
		},
		{
			name:       "MasterCard - valid prefix 55 (upper bound old range)",
			input:      "5500000000000000",
			wantValid:  true,
			wantIssuer: "Mastercard",
			wantErr:    nil,
		},
		{
			name:       "MasterCard - valid prefix 53 (middle old range)",
			input:      "5300000000000000",
			wantValid:  true,
			wantIssuer: "Mastercard",
			wantErr:    nil,
		},
		{
			name:       "MasterCard - valid prefix 2221 (lower bound new range)",
			input:      "2221000000000000",
			wantValid:  true,
			wantIssuer: "Mastercard",
			wantErr:    nil,
		},
		{
			name:       "MasterCard - valid prefix 2720 (upper bound new range)",
			input:      "2720000000000000",
			wantValid:  true,
			wantIssuer: "Mastercard",
			wantErr:    nil,
		},
		{
			name:       "MasterCard - valid prefix 2500 (middle new range)",
			input:      "2500000000000000",
			wantValid:  true,
			wantIssuer: "Mastercard",
			wantErr:    nil,
		},
		{
			name:       "MasterCard - invalid prefix 2220 (below new range)",
			input:      "2220000000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
		{
			name:       "MasterCard - invalid prefix 2721 (above new range)",
			input:      "2721000000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
		{
			name:       "MasterCard - invalid prefix 50 (below old range)",
			input:      "5000000000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
		{
			name:       "MasterCard - invalid prefix 56 (above old range)",
			input:      "5600000000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},

		// Discover Tests
		{
			name:       "Discover - valid prefix 65",
			input:      "6500000000000000",
			wantValid:  true,
			wantIssuer: "Discover",
			wantErr:    nil,
		},
		{
			name:       "Discover - valid prefix 6011",
			input:      "6011000000000000",
			wantValid:  true,
			wantIssuer: "Discover",
			wantErr:    nil,
		},
		{
			name:       "Discover - valid prefix 644 (lower bound)",
			input:      "6440000000000000",
			wantValid:  true,
			wantIssuer: "Discover",
			wantErr:    nil,
		},
		{
			name:       "Discover - valid prefix 649 (upper bound)",
			input:      "6490000000000000",
			wantValid:  true,
			wantIssuer: "Discover",
			wantErr:    nil,
		},
		{
			name:       "Discover - valid prefix 646 (middle)",
			input:      "6460000000000000",
			wantValid:  true,
			wantIssuer: "Discover",
			wantErr:    nil,
		},
		{
			name:       "Discover - valid prefix 622126 (lower bound)",
			input:      "6221260000000000",
			wantValid:  true,
			wantIssuer: "Discover",
			wantErr:    nil,
		},
		{
			name:       "Discover - valid prefix 622925 (upper bound)",
			input:      "6229250000000000",
			wantValid:  true,
			wantIssuer: "Discover",
			wantErr:    nil,
		},
		{
			name:       "Discover - valid prefix 622500 (middle)",
			input:      "6225000000000000",
			wantValid:  true,
			wantIssuer: "Discover",
			wantErr:    nil,
		},
		{
			name:       "Discover - invalid prefix 643 (below 644-649 range)",
			input:      "6430000000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
		{
			name:       "Discover - invalid prefix 650 (above 644-649 range)",
			input:      "6500000000000001",
			wantValid:  true,
			wantIssuer: "Discover",
			wantErr:    nil,
		},
		{
			name:       "Discover - invalid prefix 622125 (below 622126-622925 range)",
			input:      "6221250000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
		{
			name:       "Discover - invalid prefix 622926 (above 622126-622925 range)",
			input:      "6229260000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},

		// Edge Cases
		{
			name:       "Unknown issuer - prefix 1",
			input:      "1234567890123456",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
		{
			name:       "Unknown issuer - prefix 9",
			input:      "9234567890123456",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
		{
			name:       "Unknown issuer - prefix 0",
			input:      "0234567890123456",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
		{
			name:       "Prefix 3 but not Amex or JCB",
			input:      "3000000000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
		{
			name:       "Prefix 6 but not Discover",
			input:      "6000000000000000",
			wantValid:  false,
			wantIssuer: "",
			wantErr:    ErrIssuerNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, gotIssuer, gotErr := validateIssuerNetwork(tt.input)

			if gotValid != tt.wantValid {
				t.Errorf("ValidateIssuerNetwork() valid = %v, want %v", gotValid, tt.wantValid)
			}

			if gotIssuer != tt.wantIssuer {
				t.Errorf("ValidateIssuerNetwork() issuer = %v, want %v", gotIssuer, tt.wantIssuer)
			}

			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("ValidateIssuerNetwork() error = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func BenchmarkValidateIssuerNetwork(b *testing.B) {
	testCases := []string{
		"4532015112345678", // Visa
		"340000000000009",  // Amex
		"5100000000000000", // MasterCard
		"6011000000000000", // Discover
		"3528000000000000", // JCB
	}

	for _, tc := range testCases {
		b.Run(tc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				validateIssuerNetwork(tc)
			}
		})
	}
}
