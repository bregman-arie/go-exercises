// 1. Run `go mod init hello-world` to initialize the project (this creates a new module)
//   - This creates a file called `go.mod` that specifies the module name and the version of Go to use
// 2. Add `package main`
//   - in Go everything should be organized as package
package main

// 3. Import fmt package to use the 'print' function
import "fmt"

// 4. Add main function
//   - Go has to know where to start executing the code
//   - There can be only one function
func main() {
	fmt.Println("Hello World")
}

// If you wanted to print without a new line at the end,
// you'd use fmt.Print instead

// 5. Save and execute the program: go run main.go
