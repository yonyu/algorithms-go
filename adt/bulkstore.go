package adt

import (
	"fmt"
)

type BulkStore struct {
	capacity int
	quantity int
}

func (b BulkStore) Capacity() int {
	return b.capacity
}

func (b BulkStore) Quantity() int {
	return b.quantity
}

func (b *BulkStore) Add(amount int) (int, error) {
	if amount >= 0 {
		if b.quantity + amount > b.capacity{
			delta := b.capacity - b.quantity
			b.quantity = b.capacity
			return delta, nil
		} else {
			b.quantity += amount
			return amount, nil
		}
	} else {
		return amount, fmt.Errorf("amount is invalid")
	}
}

func (b *BulkStore) Remove(amount int) (int, error) {
	if amount >= 0 {
		if b.quantity - amount < 0 {
			delta := b.capacity - b.quantity
			b.quantity = 0
			return delta, nil
		} else {
			b.quantity -= amount
			return amount, nil
		}
	} else {
		return amount, fmt.Errorf("amount is invalid")
	}
}

func NewBulkStore(n int) *BulkStore {
	store := &BulkStore {capacity: n, quantity: 0}
	return store
}

func TestNewBulkStore() {
	var store BulkStorer = NewBulkStore(100)

	store.Add(30)

	fmt.Printf("%v\n", store)
}
