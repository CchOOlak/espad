package appErrors

import "errors"

var (
	NotFound         = errors.New("not_found")
	InvalidInput     = errors.New("invalid_input")
	Internal         = errors.New("internal")
)
