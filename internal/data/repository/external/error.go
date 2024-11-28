package external

import "errors"

var (
	BadRequestError     = errors.New("incorrect request")
	InternalServerError = errors.New("no response from API")
)
