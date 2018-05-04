package main

import (
	"fmt"
)

func main() {
	intf := []string{"eth0"}

	var idx int
	for idx = range intf {
		// find something
	}

	// Not found, do something
	if idx == len(intf) {
		// This code block will not happens
		fmt.Println("This code block will not happens")
	}
}
