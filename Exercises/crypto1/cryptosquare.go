package cryptosquare

import (
	"fmt"
	"strings"
)




// Loop through string to check for any empty spaces
// Append letters to empty variable called newstr
// Remove any other possible punctuation
// Convert all letters to lowercase

func remover(str1 string) (newstr string) {
	for i := 0; i < len(str1); i++ {
		if str1[i] != ' ' && str1[i] != ','  && str1[i] != ';' && str1[i] != ':'{
			newstr = newstr + string(str1[i])
			// fmt.Println(newstr)
			finalresult := strings.ToLower(newstr)
			fmt.Println(finalresult)
		}
	}
	return newstr
}
