package adt

// struct version of YFI
type StructYFI struct {
	value int
}

func (yfi *StructYFI) Set(inches, feet, yards int) {
	yfi.value = inches + feet * 12 + yards * 36
} 

func (yfi StructYFI) Get() (yards, feet, inches int) {
	yards = yfi.value / 36
	feet = (yfi.value - yards * 36) / 12
	inches = (yfi.value- yards * 36 - feet * 12)
	return yards, feet, inches
}

func (yfi *StructYFI) Add(other StructYFI) {
	yfi.value += other.value
}

func (yfi *StructYFI) Sub(other StructYFI) {
	yfi.value -= other.value
}