package main

import (
	"fmt"
	"sort"
) 

func main() {
	ages := make(map[string]int) //  mapping from strings to ints
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages)
	var names []string
	for name := range ages {
		fmt.Println(name)
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
		// names := make([]string, 0, len(ages))
	}
}


func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
	return false
	}
	for k, xv := range x {
	if yv, ok := y[k]; !ok || yv != xv {
	return false
	}
 }
	return true
}