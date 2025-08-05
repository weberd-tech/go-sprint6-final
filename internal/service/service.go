package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorseCode(text string) bool {
	return !strings.ContainsFunc(text, func(r rune) bool {
		return r != '.' && r != '-' && r != '/' && r != ' '
	})
}

func ConvertMorseCode(text string) string {
	var resultString string
	var resultMorse string

	if isMorseCode(text) {
		resultString = morse.ToText(text)
		return resultString
	}
	resultMorse = morse.ToMorse(text)
	return resultMorse
}
