package android

import "errors"

var (
	ErrOptionUnsupported = errors.New("option unsupported")
	ErrMethodNotFound    = errors.New("method not found")
)
