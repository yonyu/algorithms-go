package adt

type BulkStorer interface {
	Capacity() int 
	Quantity() int 
	Add(amount int) (int, error) 
	Remove(amount int) (int, error) 
}