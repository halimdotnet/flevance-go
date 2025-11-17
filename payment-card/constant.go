package pcard

import "errors"

var (
	ErrInvalidTotalDigits  = errors.New("invalid total digits")
	ErrInvalidExpiryFormat = errors.New("invalid expiry format")
	ErrInvalidExpiryDate   = errors.New("invalid expiry date")
	ErrOverExpiry          = errors.New("over year of expiry")
	ErrInvalidPaymentCard  = errors.New("invalid payment card")
)

const (
	MIN_DIGITS = 13
	MAX_DIGITS = 19
	MIN_YEAR   = 1997
	MAX_YEAR   = 20
)
