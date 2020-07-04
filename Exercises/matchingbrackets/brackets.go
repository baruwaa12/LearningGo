package matchingbrackets

func bracketcounter(bracketstring string) (openBrackets int, closeBrackets int) {
	for _, i := range bracketstring {
		if i == '(' || i == '[' || i == '{' {
			openBrackets++
		} else {
			if i == ')' || i == ']' || i == '}'{
				closeBrackets++
			}
			
		}
	}
	return openBrackets, closeBrackets
}
