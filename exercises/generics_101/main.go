package main

// CType can be int, int32, uint, or uint32
type CType interface {
	int | int32 | uint | uint32
}

// sum1 is the normal function which returns the summation of two CType
func sum1[T CType](A T, B T) T {
	return A + B
}

func main() {

}
