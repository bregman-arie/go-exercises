# Go Exercises

:information_source: &nbsp;This repo contains questions and exercises to learn and practice Golang

:bar_chart: &nbsp;There are currently **2** exercises and questions

* Exercises
* Questions

# Golang Exercises

* [Hello World](#hello-world)

<a name="exercises-hello-world"></a>
## Hello World

|Name|Exercise|Solution|Comments|
|--------|--------|------|----|
| Hello World! | [Exercise](exercises/hello_world/exercise.md) | [Solution](exercises/hello_world/main.go) | |
| Variables & Constants | [Exercise](exercises/variables/exercise.md) | [Solution](exercises/variables/main.go) | |


# Golang Questions

## Go 101

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
<summary>What is the difference between <code>var x int = 2</code> and <code>x := 2</code>?</summary><br><b>

The result is the same, a variable with the value 2.

With <code>var x int = 2</code> we are setting the variable type to integer, while with <code>x := 2</code> we are letting Go figure out by itself the type.
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
  TODO
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
