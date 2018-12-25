Simple script for obtain zero crossing properties of signal or similar numerical slice.

**Both** function returns the places of zero crossing in data slice, both from negative and positive value.  
Slice 'f' containes places in the original slice where signal cross zero from both directions.  
Slice 'k' containes directions of zero crossing by original slice. '-1' signal crossed from positive to negative, and vice versa.

**ToPositive** function returns places of zero crossing in data slice to positive value.

**ToNegative** function returns places of zero crossing in data slice to negative value.

**Example:**  
	package main
	
	import (
		"fmt"
		"zerocross"
	)
	
	func main() {
		niz := []float64{-1, 1, 0, -1, -2, 0, -3, -5, -1, 1, 5, -6}
		zeros, directions := zerocross.Both(niz)
		fmt.Println(zeros)
		fmt.Println(directions)
	}

	Results:
	zeros: [0 1 8 10]
	directions: [1 -1 0 0 0 0 0 0 1 0 -1]
