package main

import (
	"cheyne.nz/unscramble/pkg/bag/preset"
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	mybag := preset.NewDefaultTileSet()

	message := fmt.Sprintf("My bag contains a total of %d tiles of %d unique tile types.", mybag.Count(), mybag.UniqueCount())
	fmt.Println(message)
}
