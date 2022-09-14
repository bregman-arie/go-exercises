// 1.
package main

import "fmt"

func main() {
	// Name doesn't change (usually) that's why we set it as a constant
	const name = "Leo"
	var age = "25"
	fmt.Println("Hello, my name is", name, "and I'm", age, "years old")
}

/* 2. The following program will fail simply because the
   'objective' variable is not used. In Go, you must use
   every defined variable

package main

import "fmt"

func main() {
	fmt.Println("Skynet Beta Testing")

	var objective="Protect Humanity"
}

Some will say it fails because Skynet wasn't designed
to protect humanity...it's quite good argument.
*/

/* 3. The following program will fail because squares is
      a constant. That means you shouldn't change its value
	  once declared.
*/
