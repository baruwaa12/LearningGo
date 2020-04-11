package bob

import "strings"

func hasLetters(s string) bool {
	
	for _, index := range strings.ToLower(s) {
		if index <= 'z' && 1 >= 'a' {
			return true
		}
	}
	return false
}

// Hey replies to Bob's Partner
func Hey(remark string) string {

	switch remark = strings.Trim(remark, "\t\n\r "); {
	case strings.ToUpper(remark) == remark && strings.HasSuffix(remark, "?") && hasLetters(remark):
		return "Calm down, I know what I'm doing!"
	case strings.ToUpper(remark) == remark && hasLetters(remark):
		return "Whoa, chill out!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}