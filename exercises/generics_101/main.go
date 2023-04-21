package main

import "fmt"

// CType can be int, int32, uint, or uint32
type CType interface {
	int | int32 | uint | uint32
}

// sum1 is the normal function which returns the summation of two CType
func sum1[T CType](A T, B T) T {
	return A + B
}

// sum2 is the modified version of sum1 which gets variable number of CType
func sum2[T CType](Numbers ...T) T {
	var res T

	for _, num := range Numbers {
		res += num
	}

	return res
}

func main() {
	// create a list of uint
	numbers := []uint{1, 2, 3, 4, 5}

	// use our sum2 function to get the summation
	fmt.Println(sum2(numbers...))
}
