package usenet

import "errors"

var (
	ErrCorruptedNzb = errors.New("corrupted nzb")
	ErrIgnoreFile   = errors.New("ignore file")
)
