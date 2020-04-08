// This package performs celsius and fahrenheit temperature computations

package tempconv
import "fmt"


type Celsuis float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsuis = -273.15
    FreezingC     Celsuis = 0
    BoilingC      Celsuis = 100
)

func CToF(c Celsuis) Fahrenheit { 
    return Fahrenheit(c*9/5 + 32) 
}

func FToC(f Fahrenheit) Celsuis{
     return Celsius((f -32) * 5 / 9) 
}

fmt.Printf("%g\n", BoilingC-FreezingC) // "100"
boilingF := CToF(BoilingC)
fmt.Printf("%g\n", boilingF-CToF9(FreezingC)) // "180"
fmt.Printf("%g\n", boilingF-FreezingC) // compile errorL type mismatch

var c Celsuis
var f Fahrenheit
fmt.Println(c == 0)
fmt.Println(f >= 0)
fmt.Println(c == f)
fmt.Println(c == Celsius(f))

func (c Celsius) String() string { return fmt.Sprintf("%g degrees C", c) }
func (f Fahrenheit) String() string { return }

