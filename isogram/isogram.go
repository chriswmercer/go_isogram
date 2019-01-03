//Package isogram contains a function to check if a string is an isogram
package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram will return TRUE if the string is an isogram, otherwise false
func IsIsogram(input string) bool {
	if input == "" {
		return true
	}

	input = strings.ToUpper(input)

	for i, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}

		checkString := input[i+1:]

		if strings.Contains(checkString, string(r)) {
			return false
		}
	}

	return true
}