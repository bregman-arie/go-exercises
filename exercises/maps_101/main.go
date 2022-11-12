package main

import "fmt"

func main() {
	var profile = make(map[string]string)

	profile["firstName"] = "Alexander"
	profile["lastName"] = "Supertramp"
	profile["age"] = strconv.FormatUnit(unit64(24), 10)

	fmt.Print(profile)
}
