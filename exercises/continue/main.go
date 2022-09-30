package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i < 100; i++ {
		if i == 50 {
			continue
		}

		sum += i
	}

	fmt.Println("Value: %v", sum)
}

// The value at the end is 4900
