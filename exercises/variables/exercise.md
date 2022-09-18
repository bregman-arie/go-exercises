# Variables

## Objectives

1. Write a Go program that will store the values of your name and age in variables (or constants), and will print the following sentence using the variables you defined: `My name is ___ and I'm ___ years old`
   
2. Explain why the following program will fail?

```Go
package main

import "fmt"

func main() {
    fmt.Println("Skynet Beta Testing")

    var objective="Defend Humanity"
}
```

3. What's the output of the following program?

```Go
package main

import "fmt"

func main() {
    const squares=2
    var circles=0
    
    fmt.Println("Squares:", squares)
    fmt.Println("Circles:", circles)

    squares=1
    circles=7

    fmt.Println("Squares:", squares)
    fmt.Println("Circles:", circles)
}
```

## Solution

Click [here](main.go) to view the solution