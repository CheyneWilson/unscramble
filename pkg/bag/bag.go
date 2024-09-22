package bag

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"maps"
	"math/rand"
	"os"
)

type Bag[T comparable] struct {
	m map[T]int
}
type Nothing struct{}
type Set[T comparable] map[T]Nothing

// New
// - Create a new Bag filled with the items
func New[T comparable](items ...T) *Bag[T] {
	b := new(Bag[T])
	(*b).m = make(map[T]int)
	for _, item := range items {
		b.Add(item)
	}
	return b
}

// Add
// - Add an item to the bag
func (b *Bag[T]) Add(item T) {
	if b.Has(item) {
		(*b).m[item] += 1
	} else {
		(*b).m[item] = 1
	}
}

// UniqueItems
// - Return a set of all the unique items
func (b *Bag[T]) UniqueItems() Set[T] {
	keys := make(Set[T], len((*b).m))
	for k := range maps.Keys((*b).m) {
		keys[k] = Nothing{}
	}
	return keys
}

// FindRandom
// - Return an item at random from the bag. The contents of the bag is unchanged
func (b *Bag[T]) FindRandom() (T, error) {
	itemCount := (*b).Count()
	if itemCount == 0 {
		var noop T
		return noop, fmt.Errorf("no items in bag")
	}
	items := (*b).toArray()

	item := items[rand.Intn(itemCount)]
	return item, nil
}

// TakeRandom
// - Return an item at random from the bag. The item is removed from the bag
func (b *Bag[T]) TakeRandom() (T, error) {
	item, err := (*b).FindRandom()
	if err == nil {
		(*b).Remove(item)
	}
	return item, err
}

// toArray
// - transform the internal map into an array. This is a helper function for FindRandom
func (b *Bag[T]) toArray() []T {
	// TODO: If performance is an issue, review the underling map data-structure because FindRandom doesn't map
	//       efficiently to it. Also, revisit Count()
	items := make([]T, (*b).Count())
	i := 0
	for v, itemCount := range (*b).m {
		for j := 0; j < itemCount; j++ {
			items[i] = v
			i++
		}
	}
	return items
}

// Remove
//   - Remove a single item from the bag (if present).
//     Returns true if an item was removed and false otherwise.
func (b *Bag[T]) Remove(item T) bool {
	if b.Has(item) {
		if count := (*b).m[item]; count <= 1 {
			delete((*b).m, item)
		} else {
			(*b).m[item] -= 1
		}
		return true
	} else {
		return false
	}
}

// Count
// - Return the total number of items in a bag
func (b *Bag[T]) Count() int {
	var total int = 0
	for _, quantity := range (*b).m {
		total += quantity
	}
	return total
}

// CountOf
// Return the quantity of an item in the bag
func (b *Bag[T]) CountOf(item T) int {
	return (*b).m[item]
}

// UniqueCount
// - return the total number of unique items in a bag
func (b *Bag[T]) UniqueCount() int {
	return int(len((*b).m))
}

// Has
// - return true if an item is found in a bag
func (b *Bag[T]) Has(item T) bool {
	_, present := (*b).m[item]
	return present
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ToJSON
// outputs the contents of a bag to a JSON string
func (b *Bag[T]) ToJSON() ([]byte, error) {
	return json.MarshalIndent(b.m, "", "    ")
}

// ExportJson
// - write the contents of a bag to a JSON file
func (b *Bag[T]) ExportJson(path string) {

	f, err := os.Create(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			slog.Error(err.Error())
		}
	}(f)

	check(err)

	data, err := (*b).ToJSON()
	check(err)
	_, err = f.Write(data)
	check(err)
}

// ImportJson
// - read the contents of a jsonfile to construct a new bag
func ImportJson(path string) (*Bag[string], error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		slog.Error(err.Error())
	}

	b := New[string]()
	check(err)

	err = json.Unmarshal(contents, &b.m)
	if err != nil {
		slog.Error(err.Error())
	}
	return b, err
}
