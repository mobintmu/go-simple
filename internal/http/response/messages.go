package response

import "errors"

var (
	ErrInvalidID = errors.New("invalid ID")
	ErrNotFound  = errors.New("not found")
)
