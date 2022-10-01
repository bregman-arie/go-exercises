# Conditionals

## Objectives

1. Generate a random number between 1 and 100
   1. If the number is higher than 50 print "It's closer to 100"
   2. If the number is lower than 50 print "It's closer to 0"
   3. Print the generated random number
2. Modify the previous code to print "It's 50!" if the random number is 50
3. Modify the conditional in the code you previously written to check not only if the number is higher than 50 but also if's it's even. If it's even and higher than 50, print "It's closer to 100 and it's even!"


## Solution

1. 

```Go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    var randNum int = rand.Intn(100)

    if randNum > 50 {
        fmt.Println("It's closer to 100")
    } else {
        fmt.Println("It's closer to 0")
    }                                      
    fmt.Printf("Generated number: %v\n", randNum)
}
```

2. 

```Go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    var randNum int = rand.Intn(100)

    if randNum > 50 {
        fmt.Println("It's closer to 100")
    } else if randNum < 50 {
        fmt.Println("It's closer to 0")
    } else {
        fmt.Println("It's 50!")
    }
                             
    fmt.Printf("Generated number: %v\n", randNum)
}
```

3. 

```Go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    var randNum int = rand.Intn(100)

    if randNum > 50 && randNum%2 == 0 { 
        fmt.Println("It's closer to 100 and even")
    } else if randNum < 50 {
        fmt.Println("It's closer to 0")
    } else if randNum == 50 {                     
        fmt.Println("It's 50!")
    }
    fmt.Printf("Generated number: %v\n", randNum)
}
```