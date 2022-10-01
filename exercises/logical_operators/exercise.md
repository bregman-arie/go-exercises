# Logical Operators

## Objectives

1. What is the output of the following code:

```Go
package main

import "fmt"

func main() {
    x := 2017
    result1 := x > 50 && x < 2020
    result2 := x > 50 && x%2 == 0
    result3 := x%2 == 1 && x+3 == 2025
    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    x += 5
    result1 = x > 50 && x < 3000
    result2 = x > 50 && x < 3000
    result3 = x > 50 && x < 3000   
    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
}
```

2. What is the output of the following code:

```Go
package main

import "fmt"

func main() {
    x := 2017
    result1 := x > 50 || x < 2020
    result2 := x > 50 || x%2 == 0
    result3 := x%2 == 1 || x+3 == 2025
    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    x += 5
    result1 = x > 50 || x < 3000
    result2 = x > 50 || x < 3000
    result3 = x > 50 || x < 3000
    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
}
```

## Solution

Click [here](solution.md) to view the solution