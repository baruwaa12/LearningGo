package luhn

import (
	"strings"
	"unicode"
)

// Valid determine whether or not a number is valid per the luhn formula

func valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) <= 1 {
		return false
	}
	var sum int

	for i := len(runes) - 1; i >= 0; i-- {
		if !unicode.IsDigit(runes[i]) {
			return false
		}

		num := int(runes[i] - '0')
		if (len(runes)-i)%2 == 0 {
			return false
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		sum += num
	}

	return sum%10 == 0
}
