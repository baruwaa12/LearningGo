//Package diffofsquares provides a good math function
package diffofsquares

//SquareOfSum returns the square of sum of first n numbers
func SquareOfSum(num int) int {
	return num * (num + 1) * num * (num + 1) / 4
}

//SumOfSquares retuns the sum of squares fo first n numbers
func SumOfSquares(num int) int {
	return num * (num + 1) * (2*num + 1) / 6
}

//Difference returns SquareOfSum-SumOfSquares
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)  

}