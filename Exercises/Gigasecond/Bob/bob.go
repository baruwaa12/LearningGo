 package bob
 
 import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	// If you address him without actually saying anything
	if remark == "" {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(remark, "?")
	isUppercase := strings.ToUpper(remark) == remark
	hasNoLetters, _ := regexp.MatchString("^[^a-zA-Z]+$", remark)

	if isQuestion && isUppercase && !hasNoLetters {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion {
		return "Sure."
	}

	if hasNoLetters {
		return "Whatever."
	}

	if isUppercase {
		return "Whoa, chill out!"
	}

	// To anything else
	return "Whatever."
}