package FloatingPoint

for x := 0; x < 8; x++ {
	fmt.Printf("x = %d eA = %8.3f\n", x, math.Exp(float64(x)))
	}

var z float64
fmt.Println(z, -1, 1/z, -1/z, z/z) "0 -0 +Inf -Inf NaN"