# Loops 102

## Objectives

1. Define the following list `{'red', 'green', 'blue'}`
   1. Iterate over the list and print each item on its own line
2. Define a variable with a value of 1 and using a for loop add 3 3 to the variable's value, every loop. Make sure to define the loop in a way where it stops once the variable value is bigger than 100
   1. Print the final value

## Solution

1. 

```Go
package main

import "fmt"

func main() {
    var rgb = [3]string{"red", "green", "blue"}

    for _, color := range rgb {
        fmt.Println(color)
    }
}
```

2. 

```Go
package main

import "fmt"

func main() {
    i := 1

    for i < 101 {
        i += 3
    }
  
    fmt.Println("Value: %v", i)          
}
```