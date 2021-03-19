package errors

import "errors"

var (
	ErrEmptyFromLink     = errors.New("empty from link")
	ErrEmptyToLink       = errors.New("empty to link")
	ErrFromAlreadyExist  = errors.New("from already exist")
	ErrCantInsertNewData = errors.New("cant insert new data")
)
