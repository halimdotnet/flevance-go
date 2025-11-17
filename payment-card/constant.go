package pcard

import "errors"

var (
	ErrInvalidTotalDigits = errors.New("invalid total digits")
	ErrInvalidExpiryDate  = errors.New("invalid expiry date")
	ErrIssuerNotFound     = errors.New("issuer not found")
	ErrInvalidNumber      = errors.New("invalid payment number")
)

const (
	MIN_DIGITS = 13
	MAX_DIGITS = 19
	MIN_YEAR   = 1997
	MAX_YEAR   = 20
)
