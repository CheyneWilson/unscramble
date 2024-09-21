package preset

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"unicode"
)

func TestCountOfDefaultTileSet(t *testing.T) {
	bag := NewDefaultTileSet()

	if count := bag.Count(); count != 100 {
		t.Fatalf(`Bag should have 100 items, got '%d'.`, count)
	}
	if unique := bag.UniqueCount(); unique != 27 {
		t.Fatalf(`Bag should have 27 unique items, got '%d'.`, unique)
	}
}

func TestCountOfItemsInDefaultTileSet(t *testing.T) {
	bag := NewDefaultTileSet()
	assert.Equal(t, 9, bag.CountOf("A"))
	assert.Equal(t, 12, bag.CountOf("E"))
	assert.Equal(t, 9, bag.CountOf("I"))
	assert.Equal(t, 8, bag.CountOf("O"))
	assert.Equal(t, 4, bag.CountOf("U"))
	assert.Equal(t, 2, bag.CountOf("Y"))

}

func TestBag_ListUniqueItemsOfDefaultTileSet(t *testing.T) {
	bag := NewDefaultTileSet()

	alphabetPlusWild := bag.UniqueItems()
	for r := 'a'; r < 'z'; r++ {
		alpha := string(unicode.ToUpper(r))
		_, present := alphabetPlusWild[alpha]
		if !present {
			t.Fatalf(`Bag is missing '%s'.`, alpha)
		}

		alpha = "_"
		_, present = alphabetPlusWild[alpha]
		if !present {
			t.Fatalf(`Bag is missing '%s'.`, alpha)
		}
	}

	for i := 0; i <= 9; i++ {
		alpha := strconv.Itoa(i)
		_, present := alphabetPlusWild[alpha]
		if present {
			t.Fatalf(`Bag should not contain any digits '%s'.`, alpha)
		}
	}
}
