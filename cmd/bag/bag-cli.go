package main

import (
	"cheyne.nz/unscramble/pkg/bag"
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	message := bag.Hello("Jupiter")
	fmt.Println(message)
}
