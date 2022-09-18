package main

import "fmt"

func main() {
	var customers = 4
	var packagesToDeliver = 100
	fmt.Println("I have to deliver", packagesToDeliver, "packages")

	var deliveredPackages = 20
	packagesToDeliver -= deliveredPackages
	fmt.Printf("I have delivered %v packages\n", deliveredPackages) // Note that from this point on we are using printf and not println
	fmt.Printf("Remaining packages to deliver: %v\n", packagesToDeliver)

	var packagesPerCustomer = packagesToDeliver / customers
	fmt.Printf("Packages are going to be distributed equally between %v customers. This means %v packages per customer",
		customers, packagesPerCustomer)
}
