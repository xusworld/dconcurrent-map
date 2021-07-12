package cmap

import "errors"

var (
	errKeyNotFound    = errors.New("key not found")
	errInvalidKeyType = errors.New("invalid key type")
)
