package secret

import (
	"fmt"
)

var actions = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
	"reverse",
}

// Handshake converts a code uint into a series of secret actions
func Handshake(code uint) (seriesOfActionis []string) {
	binary :=fmt.Sprintf("%0b", code)
	res := []string{}
	reverse := false


	for i := len(binary) - 1; i > -1; i-- {
		r := binary[i]
		if r != '1' {
			continue
		}

		action := actions[len(binary)-1-i]
		if action == "reverse" {
			reverse = true
			continue
		}
		res = append(res, action)
	}

	// reverse the result slice
	if reverse {
		for a, b := 0, len(res)-1; a < b; a, b = a+1, b-1 {
			res[a], res[b] = res[b], res[a]
		}	
	}	
	return res	
}