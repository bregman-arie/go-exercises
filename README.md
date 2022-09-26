# Go Exercises

:information_source: &nbsp;This repo contains questions and exercises to learn and practice Golang

:bar_chart: &nbsp;There are currently **20** exercises and questions

- [Go Exercises](#go-exercises)
  - [Exercises](#exercises)
    - [Hello World](#hello-world)
    - [Strings](#strings)
    - [Arrays](#arrays)
    - [Loops](#loops)
  - [Questions](#questions)
    - [Go 101](#go-101)
    - [Data Types](#data-types)
    - [User Input](#user-input)
    - [Functions](#functions)
  - [Projects](#projects)

## Exercises

### Hello World

|Name|Exercise|Solution|Comments|
|--------|--------|------|----|
| Hello World! | [Exercise](exercises/hello_world/exercise.md) | [Solution](exercises/hello_world/main.go) | |
| Variables & Constants | [Exercise](exercises/variables/exercise.md) | [Solution](exercises/variables/main.go) | |
| Arithmetic Operators | [Exercise](exercises/arithmetic_operators/exercise.md) | [Solution](exercises/arithmetic_operators/main.go) | |
| Data Types | [Exercise](exercises/data_types/exercise.md) | [Solution](exercises/data_types/solution.md) | |
| User Input | [Exercise](exercises/user_input/exercise.md) | [Solution](exercises/user_input/main.go) | |
| Packages | [Exercise](exercises/packages/exercise.md) | [Solution](exercises/packages/main.go) | 


### Strings

|Name|Exercise|Solution|Comments|
|--------|--------|------|----|
| Split Strings | [Exercise](exercises/split_strings/exercise.md) | [Solution](exercises/split_strings/solution.md) | |

### Arrays

|Name|Exercise|Solution|Comments|
|--------|--------|------|----|
| Arrays 101 | [Exercise](exercises/arrays_101/exercise.md) | [Solution](exercises/arrays_101/solution.md) | |
| Slices 101 | [Exercise](exercises/slices_101/exercise.md) | [Solution](exercises/slices_101/solution.md) | |

### Loops

|Name|Exercise|Solution|Comments|
|--------|--------|------|----|
| Loops 101 | [Exercise](exercises/loops_101/exercise.md) | [Solution](exercises/loops_101/main.go) | |


## Questions

### Go 101

<details>
<summary>What are some characteristics of the Go programming language?</summary><br><b>

  * Strong and static typing - the type of the variables can't be changed over time and they have to be defined at compile time
  * Simplicity
  * Fast compile times
  * Built-in concurrency
  * Garbage collected
  * Platform independent
  * Compile to standalone binary - anything you need to run your app will be compiled into one binary. Very useful for version management in run-time.
  * Big community
</b></details>

<details>
<summary>True or False?

  * In Go we can redeclare variables
  * Once a variable declared, we must use it
</summary>

* False
* True
</b></details>

<details>
<summary>What libraries of Go have you used?</summary><br><b>

This should be answered based on your usage but some examples are:

  * fmt - formatted I/O
  * math - random numbers, arithmetic operations like square, etc.
  * time - for any time related operation

</b></details>

<details>
<summary>Why some functions and packages in Go begin with a Capital letter?</summary><br><b>

Any exported variable, function, ... begins with a capital letter. In fact when using a package, you can use only the things the package exports for you and others to use.
</b></details>

### Data Types

<details>
<summary>What is the result of the following program?

```
package main

import "fmt"

func main() {
    var userName
    userName = "user"
    fmt.Println(userName)
}
```
</summary><br><b>

Error. The type userName is not defined. It has to be declared or the value assignment should happen at declaration.

So both `var userName = user` and `var userName string` are valid.

</b></details>

<details>
<summary>What is the result of the following program?

```Go
package main
 
import "fmt"
 
func main() {
    var x uint
    x = -3
    // comment!
    fmt.Println(x)
}
```
</summary><br><b>

Error. When `x` is declared as `uint` it means you can't put any negative values. But because we did put a negative value (`-3`) it will fail to compile it.
</b></details>

<details>
<summary>What is the problem with the following block of code? How to fix it?

```
func main() {
    var x float32 = 13.5
    var y int
    y = x
}
```
</summary><br><b>
</b></details>

<details>
<summary>What is the difference between <code>var x int = 2</code> and <code>x := 2</code>?</summary><br><b>

The result is the same, a variable with the value 2.

With <code>var x int = 2</code> we are setting the variable type to integer, while with <code>x := 2</code> we are letting Go figure out by itself the type.
</b></details>

<details>
<summary>The following block of code tries to convert the integer 101 to a string but instead we get "e". Why is that? How to fix it?

```
package main

import "fmt"

func main() {
    var x int = 101
    var y string
    y = string(x)
    fmt.Println(y)
}
```
</summary><br><b>

It looks what unicode value is set at 101 and uses it for converting the integer to a string.
If you want to get "101" you should use the package "strconv" and replace <code>y = string(x)</code> with <code>y = strconv.Itoa(x)</code>
</b></details>

<details>
<summary>What is wrong with the following code?:

```
package main

func main() {
    var x = 2
    var y = 3
    const someConst = x + y
}
```
</summary><br><b>
Constants in Go can only be declared using constant expressions.
But `x`, `y` and their sum is variable.
<br>
<code>const initializer x + y is not a constant</code>
</b></details>

### User Input

<details>
<summary>Fix the following program to get an input from the user regarding his name and print it

```Go
var name string

fmt.Scan(name)

fmt.Println(name)
```
</summary><br><b>

```Go
var name string

fmt.Scan(&name)

fmt.Println(name)
```

</b></details>

<details>
<summary>Why when asking for user input, we have to specify &?</summary><br><b>

Because we want to reference the memory address of the variable, this is where the user input will be stored.
</b></details>

<details>
<summary>What do we print here?

```Go
var age int = 3

fmt.Println(age)
fmt.Println(&age)
```
</summary><br><b>

The value of `age` variable and then the memory location of `age` variable

</b></details>

### Functions

<details>
<summary>Define a function that prints "Hello World"</summary><br><b>

```Go
func PrintHelloWorld() {
  fmt.Println("Hello World")
}
```
</b></details>

<details>
<summary>What the following code does?

```Go
func add(x int, y int) int {
	return x + y
}
```
</summary><br><b>

add is a function that takes two parameters (two integers) and returns their sum.
</b></details>

<details>
<summary>Which of the following functions will work just fine?

```Go
func add(x int, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}
```
</summary><br><b>

Both. If two or more parameters share the same type, you can specify it only once.
</b></details>

## Projects

TODO