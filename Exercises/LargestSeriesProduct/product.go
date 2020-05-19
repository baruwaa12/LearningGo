package main

import (
	// "bufio"
	"fmt"
	"strconv"
	// "os"
)

func collectnumber() {
	// number := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your number: ")
	// number, _ := reader.ReadString('\n')
}

func ndecider(mynumber string) {
	fmt.Println(len(mynumber))

	for i := 0; i <= len(mynumber)-3; i++ {
		a, _ := strconv.Atoi(string(mynumber[i]))
		b, _ := strconv.Atoi(string(mynumber[i + 1]))
		c, _ := strconv.Atoi(string(mynumber[i + 2]))
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println("-----------")

	}

}

func main() {
	fmt.Println("Done")
	ndecider("1023344898123")
}
