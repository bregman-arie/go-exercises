package main

import (
	"fmt"
	"strconv"
)

func main() {
	var profile = make(map[string]string)

	profile["firstName"] = "Alexander"
	profile["lastName"] = "Supertramp"
	profile["age"] = strconv.FormatUint(uint64(24), 10)

	fmt.Print(profile)
}
