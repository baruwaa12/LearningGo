package cryptosquare

import (
	"strings"
	"unicode"
)

func Get_Row_Column(text string)(int, int){
	col := 0
	row := 0

	for col * row < len(text) {
		col++
		if col * row >= len(text){
			break
		}
		row++
	}

	return col, row
}

func Encode(text string) string{
	reduce_text := ""
	text = strings.ToLower(text)

	for _, char := range text{
		if unicode.IsLetter(char) || unicode.IsNumber(char){
			reduce_text += strings.ToLower(string(char))
		}
	}

	//re := regexp.MustCompile("(?m)[a-z]|[0-9]")
	//for _, match := range re.FindAllString(text, -1) {
	//	reduce_text += match
	//}

	if reduce_text == ""{
		return ""
	}

	col, row := Get_Row_Column(reduce_text)

	for len(reduce_text) < (col * row){
		reduce_text += " "
	}

	result := ""

	for c := 0; c < col; c++{
		if result != ""{
			result += " "
		}
		start_idx := c
		for r := 0; r < row; r++{
			result += string(reduce_text[start_idx])
			start_idx += col
		}
	}
	return result
}