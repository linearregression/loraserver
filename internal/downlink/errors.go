package downlink

import "errors"

// downlink errors
var (
	ErrFPortMustNotBeZero     = errors.New("FPort must not be 0")
	ErrFPortMustBeZero        = errors.New("FPort must be 0")
	ErrNoLastRXInfoSet        = errors.New("no last RX-Info set available")
	ErrInvalidDataRate        = errors.New("invalid data-rate")
	ErrMaxPayloadSizeExceeded = errors.New("maximum payload size exceeded")
)
