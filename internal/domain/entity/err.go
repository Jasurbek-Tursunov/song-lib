package entity

import "errors"

var (
	BadRequestError = errors.New("incorrect request")
	NotFoundError   = errors.New("object not found")
	InternalError   = errors.New("no response")
)
