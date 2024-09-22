package preset

import (
	"cheyne.nz/unscramble/pkg/bag"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"unicode"
)

func TestCountOfDefaultTileSet(t *testing.T) {
	alphabetBag := NewDefaultTileSet()

	if count := alphabetBag.Count(); count != 100 {
		t.Fatalf(`Bag should have 100 items, got '%d'.`, count)
	}
	if unique := alphabetBag.UniqueCount(); unique != 27 {
		t.Fatalf(`Bag should have 27 unique items, got '%d'.`, unique)
	}
}

func TestCountOfItemsInDefaultTileSet(t *testing.T) {
	alphabetBag := NewDefaultTileSet()
	assert.Equal(t, 9, alphabetBag.CountOf("A"))
	assert.Equal(t, 12, alphabetBag.CountOf("E"))
	assert.Equal(t, 9, alphabetBag.CountOf("I"))
	assert.Equal(t, 8, alphabetBag.CountOf("O"))
	assert.Equal(t, 4, alphabetBag.CountOf("U"))
	assert.Equal(t, 2, alphabetBag.CountOf("Y"))

}

func TestBag_ListUniqueItemsOfDefaultTileSet(t *testing.T) {
	alphabetBag := NewDefaultTileSet()

	alphabetPlusWild := alphabetBag.UniqueItems()
	// The bag should contain every letter of the alphabet
	for r := 'a'; r < 'z'; r++ {
		alpha := string(unicode.ToUpper(r))
		_, present := alphabetPlusWild[alpha]
		if !present {
			t.Fatalf(`Bag is missing '%s'.`, alpha)
		}
	}

	// It should also contain an underscore
	underscore := "_"
	_, present := alphabetPlusWild[underscore]
	if !present {
		t.Fatalf(`Bag is missing '%s'.`, underscore)
	}

	// It should not contain any digits
	for i := 0; i <= 9; i++ {
		alpha := strconv.Itoa(i)
		_, present := alphabetPlusWild[alpha]
		if present {
			t.Fatalf(`Bag should not contain any digits '%s'.`, alpha)
		}
	}
}

func TestBag_FindRandom(t *testing.T) {
	alphabetBag := NewDefaultTileSet()
	initialCount := alphabetBag.Count()
	item1, err := alphabetBag.FindRandom()
	assert.Nil(t, err)
	item2, err := alphabetBag.FindRandom()
	assert.Nil(t, err)
	item3, err := alphabetBag.FindRandom()
	assert.Nil(t, err)
	item4, err := alphabetBag.FindRandom()
	assert.Nil(t, err)

	finalCount := alphabetBag.Count()
	if initialCount != finalCount {
		t.Fatalf(`The size of the bag should be unchanged. Expected size of %d, got '%d'.`, initialCount, finalCount)
	}

	if !alphabetBag.Has(item1) {
		t.Fatalf(`Bag does not contain expected %s.`, item1)
	}
	if !alphabetBag.Has(item2) {
		t.Fatalf(`Bag does not contain expected %s.`, item2)
	}
	if !alphabetBag.Has(item3) {
		t.Fatalf(`Bag does not contain expected %s.`, item3)
	}
	if !alphabetBag.Has(item4) {
		t.Fatalf(`Bag does not contain expected %s.`, item4)
	}
}

func TestBag_TakeRandom(t *testing.T) {
	alphabetBag := NewDefaultTileSet()
	referenceBag := NewDefaultTileSet()
	removedItemsBag := bag.New[string]()

	initialCount := alphabetBag.Count()

	item, err := alphabetBag.TakeRandom()
	assert.Nil(t, err)
	removedItemsBag.Add(item)

	// The -1 is to test for a dumb mistakes, e.g. it would catch if referenceBag and bag were reference to the same object.
	assert.Equal(t, referenceBag.CountOf(item)-1, alphabetBag.CountOf(item))
	assert.Equal(t, referenceBag.CountOf(item)-removedItemsBag.CountOf(item), alphabetBag.CountOf(item))

	item, err = alphabetBag.TakeRandom()
	assert.Nil(t, err)
	removedItemsBag.Add(item)
	expectedRemaining := referenceBag.CountOf(item) - removedItemsBag.CountOf(item)
	assert.Equal(t, expectedRemaining, alphabetBag.CountOf(item))

	item, err = alphabetBag.TakeRandom()
	assert.Nil(t, err)
	removedItemsBag.Add(item)
	expectedRemaining = referenceBag.CountOf(item) - removedItemsBag.CountOf(item)
	assert.Equal(t, expectedRemaining, alphabetBag.CountOf(item))

	item, err = alphabetBag.TakeRandom()
	assert.Nil(t, err)
	removedItemsBag.Add(item)
	expectedRemaining = referenceBag.CountOf(item) - removedItemsBag.CountOf(item)
	assert.Equal(t, expectedRemaining, alphabetBag.CountOf(item))

	finalCount := alphabetBag.Count()
	expectedFinalCount := initialCount - 4
	if expectedFinalCount != finalCount {
		t.Fatalf(`The size of the should 4 less. Expected size of %d, got '%d'.`, expectedFinalCount, finalCount)
	}
}
