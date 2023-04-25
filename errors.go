package serialisation

import (
	"errors"
)

var ErrInvalidFormat error = errors.New("invalid format")

var ErrConversion error = errors.New("conversion error")
