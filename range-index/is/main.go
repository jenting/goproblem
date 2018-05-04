package main

import (
	"fmt"
)

func main() {
	intf := []string{"eth0"}

	var idx int
	for idx = 0; idx < len(intf); idx++ {
		// find something
	}

	// Not found, do something
	if idx == len(intf) {
		// This code block will happens
		fmt.Println("This code block will happens")
	}
}
