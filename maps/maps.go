package main

import (
	"fmt"
) 


func main(){
	fmt.Println("start")
	name()
}

func name() {
	// var names []string
	ages := make(map[string]int)
	ages["alice"] = 31
	
	// for name := range ages{names = append(names, name)
	
	// }

	for a, b := range ages{
		fmt.Println(a, b)
	}

}

// class Ages{
// 	string[] Names
// }

// sort.Strings(names)
// for _, name := range names {
// 	fmt.Printf("%s\t%d\n", name, ages[name])
// }