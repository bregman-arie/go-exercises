# Split Strings

## Objectives

1. Define a variable with the value of this name/string: `Alexander Supertramp`
2. Using the variable you defined, print only the first name
3. Using the variable you defined, print only the surname

## Solution

```Go
package main
        
import (
    "fmt"
    "strings"
)

func main() {

    var name string = "Alexander Supertramp"
    fmt.Printf("First name: %v\n", strings.Fields(name)[0])                                                        
    fmt.Printf("Surname: %v", strings.Fields(name)[1])
}
```