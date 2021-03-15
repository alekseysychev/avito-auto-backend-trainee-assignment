package errors

import "errors"

var (
	ErrEmptyFromLink = errors.New("empty from link")
	ErrEmptyToLink   = errors.New("empty to link")
)
