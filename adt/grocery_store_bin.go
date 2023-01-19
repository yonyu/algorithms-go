package adt

import (
	"fmt"
)

type GroceryStoreBin struct {
	serial int
	BulkStore
}

func NewGroceryStoreBin(serial, n int) *GroceryStoreBin {
	var bin GroceryStoreBin = GroceryStoreBin{}
	bin.serial = serial
	bin.capacity = n
	return &bin
}

func TestNewGroceryStoreBin() {
	var bin BulkStorer = NewGroceryStoreBin(1, 100)

	bin.Add(30)

	fmt.Printf("%v\n", bin)
}