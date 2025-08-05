package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorseCode(text string) bool {
	for _, ch := range text {
		if ch != '.' && ch != '-' && ch != '/' && ch != ' ' {
			return false
		}
	}
	return true
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
