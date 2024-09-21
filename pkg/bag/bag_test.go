package bag

import (
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
