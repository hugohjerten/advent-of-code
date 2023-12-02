package utils

import (
	"unicode"
)

func IsByteInt(b byte) bool {
	return unicode.IsDigit(rune(b))
}
