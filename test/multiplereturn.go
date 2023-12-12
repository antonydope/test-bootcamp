package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	// Capitalize the input string
	s := strings.ToUpper(n)

	// Split the string into words
	names := strings.Split(s, " ")

	// Initialize an empty slice to store initials
	var initials []string

	// Extract the first letter of each word as initials
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	// Check if there are at least two initials
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	// If only one initial is available, return it and an underscore "_"
	return initials[0], "_"
}

func main() {
	// Test cases
	fn1, sn1 := getInitials("paul mudavadi")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("tom onyango")
	fmt.Println(fn2, sn2)
}
