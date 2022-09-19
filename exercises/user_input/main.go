package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Printf("Your name is %v and your age is %v", name, age)
}
