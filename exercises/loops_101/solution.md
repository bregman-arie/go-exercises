# Loops 101

## Objectives

1. Write a loop to ask the name and age from a user and print "Your name is ___ and your age is ___ years old"
   1. This should be infinite loop
2. Define the following array `"pizza"}`
   1. Append to it the following item: "hamburger"
   2. Append to it the following item: "salad"
   3. Iterate over the list and print for each item `Food: <Food name>`. Make sure to replace `<Food name>` with item from the array
3. Define the following fixed array of 3 items: `{square, circle, triangle}`
   1. Iterate over it and print for each item the following: `This is <ITEM> and its index in the array is <INDEX>`


## Solution

1. 

```Go
package main
            
import "fmt"
         
func main() {
    var name string
    var age uint
         
    for {
         
        fmt.Print("\nEnter your name: ")
        fmt.Scan(&name)
        fmt.Print("Enter your age: ")
        fmt.Scan(&age)
         
        fmt.Printf("\nYour name is %v and your age is %v years old\n", name, age)                                                                                            
    }       
}
```

2. 

```Go
package main
 
import "fmt"
 
func main() {
 
    var foods = []string{"pizza"}
                                                                                                                                                                             
    foods = append(foods, "hamburger")
    foods = append(foods, "salad")
 
    for _, food := range foods {
        fmt.Printf("Food: %v\n", food)
    }   
}
```

3. .

```Go
package main

import "fmt"

func main() {

    var shapes = [3]string{"square", "circle", "triangle"}
    
    for index, shape := range shapes {
        fmt.Printf("This is %v and its index in the array is %v\n", shape,index)                                        
    }
}
```