s := "hello world"
fmt.Println(len(s))
fmt.Println(s[0], s[7])


func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
}

