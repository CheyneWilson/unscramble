package bag

type Bag[T comparable] map[T]uint

// New
// - Create a new Bag filled with the items
func New[T comparable](items ...T) *Bag[T] {
	b := new(Bag[T])
	*b = make(map[T]uint)
	for _, item := range items {
		b.Add(item)
	}
	return b
}

// Add
// - Add an item to the bag
func (b *Bag[T]) Add(item T) {
	if b.Has(item) {
		(*b)[item] += 1
	} else {
		(*b)[item] = 1
	}
}

// Count
// - Return the total number of items in a bag
func (b *Bag[T]) Count() uint {
	var total uint = 0
	for _, quantity := range *b {
		total += quantity
	}
	return total
}

// UniqueCount
// - return the total number of unique items in a bag
func (b *Bag[T]) UniqueCount() uint {
	return uint(len(*b))
}

// Has
// - return true if an item is found in a bag
func (b *Bag[T]) Has(item T) bool {
	_, present := (*b)[item]
	return present
}
