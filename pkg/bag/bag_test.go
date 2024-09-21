package bag

import (
	"cheyne.nz/unscramble/pkg/bag/preset"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmptyBag(t *testing.T) {
	emptyBag := New[string]()
	if emptyBag.Count() != 0 {
		t.Fatal(`Empty bag should have zero items.`)
	}

	if emptyBag.UniqueCount() != 0 {
		t.Fatal(`Empty bag should have zero unique items.`)
	}
}

func TestBagWithUniqueLetters(t *testing.T) {
	tiles := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	bag := New(tiles...)
	if count := bag.Count(); count != 8 {
		t.Fatalf(`Bag should have 8 items, got '%d'.`, count)
	}

	if unique := bag.UniqueCount(); unique != 8 {
		t.Fatalf(`Bag should have 8 unique items, got '%d'.`, unique)
	}
}

func TestBagWithDuplicateLetters(t *testing.T) {
	tiles := []string{"a", "b", "c", "c", "c", "c", "d", "e", "e", "f", "g", "h", "h", "h"}
	bag := New(tiles...)
	if count := bag.Count(); count != 14 {
		t.Fatalf(`Bag should have 14 items, got '%d'.`, count)
	}

	if unique := bag.UniqueCount(); unique != 8 {
		t.Fatalf(`Bag should have 8 unique items, got '%d'.`, unique)
	}
}

func TestAddToEmptyBag(t *testing.T) {
	bag := New[string]()
	if bag.Count() != 0 {
		t.Fatal(`Empty bag should have zero items.`)
	}

	if bag.UniqueCount() != 0 {
		t.Fatal(`Empty bag should have zero unique items.`)
	}

	bag.Add("a")
	if count := bag.Count(); count != 1 {
		t.Fatalf(`Bag should have 1 items, got '%d'.`, count)
	}

	if unique := bag.UniqueCount(); unique != 1 {
		t.Fatalf(`Bag should have 1 unique items, got '%d'.`, unique)
	}

	bag.Add("a")
	bag.Add("a")
	bag.Add("a")
	if count := bag.Count(); count != 4 {
		t.Fatalf(`Bag should have 4 items, got '%d'.`, count)
	}

	if unique := bag.UniqueCount(); unique != 1 {
		t.Fatalf(`Bag should have 1 unique items, got '%d'.`, unique)
	}

	bag.Add("b")
	bag.Add("c")
	bag.Add("d")
	if count := bag.Count(); count != 7 {
		t.Fatalf(`Bag should have 7 items, got '%d'.`, count)
	}

	if unique := bag.UniqueCount(); unique != 4 {
		t.Fatalf(`Bag should have 4 unique items, got '%d'.`, unique)
	}
}

func TestBag_FindRandom(t *testing.T) {
	bag := preset.NewDefaultTileSet()
	initialCount := bag.Count()
	item1, err := bag.FindRandom()
	assert.Nil(t, err)
	item2, err := bag.FindRandom()
	assert.Nil(t, err)
	item3, err := bag.FindRandom()
	assert.Nil(t, err)
	item4, err := bag.FindRandom()
	assert.Nil(t, err)

	finalCount := bag.Count()
	if initialCount != finalCount {
		t.Fatalf(`The size of the bag should be unchanged. Expected size of %d, got '%d'.`, initialCount, finalCount)
	}

	if !bag.Has(item1) {
		t.Fatalf(`Bag does not contain expected %s.`, item1)
	}
	if !bag.Has(item2) {
		t.Fatalf(`Bag does not contain expected %s.`, item2)
	}
	if !bag.Has(item3) {
		t.Fatalf(`Bag does not contain expected %s.`, item3)
	}
	if !bag.Has(item4) {
		t.Fatalf(`Bag does not contain expected %s.`, item4)
	}
}

func TestBag_TakeRandom(t *testing.T) {
	bag := preset.NewDefaultTileSet()
	referenceBag := preset.NewDefaultTileSet()
	removedItemsBag := New[string]()

	initialCount := bag.Count()

	item, err := bag.TakeRandom()
	assert.Nil(t, err)
	removedItemsBag.Add(item)

	// The -1 is to test for a dumb mistakes, e.g. it would catch if referenceBag and bag were reference to the same object.
	assert.Equal(t, referenceBag.CountOf(item)-1, bag.CountOf(item))
	assert.Equal(t, referenceBag.CountOf(item)-removedItemsBag.CountOf(item), bag.CountOf(item))

	item, err = bag.TakeRandom()
	assert.Nil(t, err)
	removedItemsBag.Add(item)
	expectedRemaining := referenceBag.CountOf(item) - removedItemsBag.CountOf(item)
	assert.Equal(t, expectedRemaining, bag.CountOf(item))

	item, err = bag.TakeRandom()
	assert.Nil(t, err)
	removedItemsBag.Add(item)
	expectedRemaining = referenceBag.CountOf(item) - removedItemsBag.CountOf(item)
	assert.Equal(t, expectedRemaining, bag.CountOf(item))

	item, err = bag.TakeRandom()
	assert.Nil(t, err)
	removedItemsBag.Add(item)
	expectedRemaining = referenceBag.CountOf(item) - removedItemsBag.CountOf(item)
	assert.Equal(t, expectedRemaining, bag.CountOf(item))

	finalCount := bag.Count()
	expectedFinalCount := initialCount - 4
	if expectedFinalCount != finalCount {
		t.Fatalf(`The size of the should 4 less. Expected size of %d, got '%d'.`, expectedFinalCount, finalCount)
	}
}
