package palindrome

import (
	"fmt"
	"math"
	"sync"
)

// Product structure
type Product struct {
	Product        int
	Factorizations [][2]int
}

func isPalindrome(p int) bool {
	if p > -10 && p < 10 {
		return true
	}

	r := 0
	for n := p; n > 0; n /= 10 {
		r = r*10 + n%10
	}
	return p == r
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = fmt.Errorf("fmin > fmax")
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		min := math.MaxInt64
		for i := fmin; i <= fmax; i++ {
			for j := i; j <= fmax; j++ {
				product := i * j
				switch isPalindrome(product) {
				case product <= min:
					switch {
					case product == pmin.Product:
						pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
					case product < min:
						min = product
						pmin = Product{Product: product, Factorizations: [][2]int{{i, j}}}
					}
				}
			}
		}
		wg.Done()
	}()

	go func() {
		max := math.MinInt64
		for i := fmax; i >= fmin; i-- {
			for j := i; j >= fmin; j-- {
				product := i * j
				switch isPalindrome(product) {
				case product >= max:
					switch {
					case product == pmax.Product:
						pmax.Factorizations = append(pmax.Factorizations, [2]int{j, i})
					case product > max:
						max = product
						pmax = Product{Product: product, Factorizations: [][2]int{{j, i}}}
					}
				}
			}
		}
		wg.Done()
	}()

	wg.Wait()

	if pmin.Product == 0 && pmax.Product == 0 {
		err = fmt.Errorf("no palindromes")
	}
	return
}