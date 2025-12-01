package game

import "errors"

var (
	ErrFontNotFound    = errors.New("font file not found")
	ErrFontParseFailed = errors.New("font file could not be parsed")
)
