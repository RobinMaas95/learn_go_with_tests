package iteration

import "strings"

func Repeat(character string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return
}

func RepeatUsingStrings(character string, repeatCount int) (repeated string) {
	return strings.Repeat(character, repeatCount)
}
