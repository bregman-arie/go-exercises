# Data Types

## Objectives

1. What is the result of the following program? Why?

```Go
package main

import "fmt"

func main() {
    var userName
    userName = "user"
    fmt.Println(userName)
}
```

2. What is the result of the following program? Why?

```Go
package main

import "fmt"

func main() {
    var userName = 2
    fmt.Println(userName)
}
```

3. Fix the following program by modifying one of the lines (but not adding or removing lines)

```Go
package main

import "fmt"

func main() {
    var userName
    userName = "user"
    fmt.Println(userName)
}
```

4. Modify the following program to print the types of the variables

```Go
package main

import "fmt"

func main() {
    var food = "Pizza"
    var slices = 4
    var pineappleOnPizza = True
}
```

## Solution

1. Error. Go is statically typed and you need to tell Go Compiler the data type when doing variable declaration.

2. It will print 2. The reason it doesn't fail (even though you didn't declare the type) is due to Go performing type inference where it infers the type from the assigned value

3. Declare the type of userName

```Go
package main

import "fmt"

func main() {
    var userName string
    userName = "user"
    fmt.Println(userName)
}
```

4. 

```Go
package main

import "fmt"

func main() {
    var food = "Pizza"
    var slices = 4
    var pineappleOnPizza = true

    fmt.Printf("food is %T\nslices is %T\npineappleOnPizza is %T", food, slices, pineappleOnPizza)
}
```