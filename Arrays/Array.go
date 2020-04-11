package array

import (
	"fmt"
)



func test(){
	var a [3]int // array of 3 integers
	fmt.Println(a[0]) // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// 
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)

	
		a := [2]int{1, 2}
		b := [...]int{1, 2}
		c := [2]int{1, 3}
		fmt.Println(a == b, a == c, b == c) // "true false false"
		d := [3]int{1, 2}
		fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
	}
}


func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}


/// Append Function

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z=x[:zlen]
		} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
		zcap = 2 * len(x)
		}
		z=make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
		}
		z[len(x)] = y
		return z
	} 
}