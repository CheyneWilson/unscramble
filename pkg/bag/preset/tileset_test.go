package preset

import (
	"fmt"
	"testing"
)

func TestCountOfDefaultTileSet(t *testing.T) {
	bag := NewDefaultTileSet()

	if count := bag.Count(); count != 100 {
		t.Fatalf(`Bag should have 100 items, got '%d'.`, count)
	}
	if unique := bag.UniqueCount(); unique != 27 {
		fmt.Print(bag)
		t.Fatalf(`Bag should have 27 unique items, got '%d'.`, unique)

	}

}
