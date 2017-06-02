package ipip

import (
	"errors"
)

var (
	ErrInvalidIp  = errors.New("invalid ip")
	ErrIpNotFound = errors.New("ip not found")
)
