package constant

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrNotFound         = errors.New("not found")
)

var HttpErrors = map[error]int{
	ErrInvalidParameter: http.StatusBadRequest,
	ErrNotFound:         http.StatusNotFound,
}
