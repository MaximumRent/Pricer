package util

import (
	"strings"
	"unicode"
	"bytes"
)

func IsNonTrashString(line string) (bool) {
	return len(strings.TrimSpace(line)) > 0
}

func PickoutLettersFromString(line string) (string) {
	var buffer bytes.Buffer
	for _, char := range line {
		if !unicode.IsSpace(char) {
			buffer.WriteRune(char)
		}
	}
	return buffer.String()
}