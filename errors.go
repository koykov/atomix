package atomix

import "errors"

var (
	ErrUnsupportedOrder = errors.New("unsupported memory order")
)
