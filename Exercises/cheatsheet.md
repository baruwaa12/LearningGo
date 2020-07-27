## Creating a 2D array of any type with any size


make([]string, 1) - This example will make an array of strings with size 1

To append slices and arrays together use the append function.
Using the triple dots will allow the array to be joined together in reverse.

append(slice1, slice2)

  
## String Interpolation

For normal strings use %s
For type format use %T
For boolean format use %t
For flags use %#

## The use of deep equal

Using deep equal is only required when comparing two different array contents.

## Creating a 2d array 

make([][]int, rows) - This is a 2d array with each a number of rows.

## Typical Structure for a a test

var testCases = []struct {
	
    Whatever values you expect to return here.

    noun string
    
}

## Looping through a slice/array

inner := make([]int, column)
for j := 0; j < column; j++


