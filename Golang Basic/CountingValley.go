package main

import (
	"fmt"
	"strings"
)

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	// path = "UDDDUDUU"
	// steps = 8
	var output int32
	s := strings.Split(path, "")
	if s[0] == "U" {
		output = 1
	} else {
		output = 2
	}
	for i := 0; i < len(path); i++ {
		if s[i] == "U" {
			output += 1
		} else {
			output -= 1
		}

	}
	return int32(output)
}

func main() {

	fmt.Println(countingValleys(8, "DDUUDDUDUUUD"))
}
