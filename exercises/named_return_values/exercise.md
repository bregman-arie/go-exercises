# Names Return Values

## Objectives

1. What is the result of the following code?

```Go
package main
        
import "fmt"
        
func add(x, y int) {
    var z int = x + y                     
    return z
}
       
func main() {
    fmt.Print(add(3, 2))
}
```

2. Fix the code above

3. Will the following code work properly? Explain

```Go
package main
        
import "fmt"
        
func add(x, y int) (z int) {
    z = x + y                     
    return
}
       
func main() {
    fmt.Print(add(3, 2))
}
```

## Solution

Click [here](solution.md) to view the solution