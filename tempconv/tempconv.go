package tempconv

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr), "cf: %s\n",
			    f, tempconv.FToC(f), c, tempconv.CToF(c))
		}
	}
