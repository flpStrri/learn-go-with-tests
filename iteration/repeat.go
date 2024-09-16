package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	// var repeated_character string

	// for i := 0; i < repeatCount; i++ {
	// 	repeated_character += character
	// }
	// return repeated_character
	return strings.Repeat(character, repeatCount)
}
