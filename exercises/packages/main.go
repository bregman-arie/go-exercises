package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("The time now is:", time.Now())
	fmt.Println("Random number:", rand.Intn(100))
	fmt.Printf("The square of 9: %g\n", math.Sqrt(9))
}
