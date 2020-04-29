package twofer

import (
	"fmt"
	"strings"
	)

// ShareWith returns a string based on the input.
func ShareWith(name string) string {
	if strings.TrimSpace(name) == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}