package adt

// ADT type YFI:
// Converting between a YFI value (a distance in inch) and y yards f feet i inches
type YFI int

// The carrier set of the YFI ADT is the set of all triples (y, f, i).
// 
// The method set of the YFI ADT includes the following operations besides
// the operations of the built-in type (int) on which this type is based.

// Return the YFI value m as a triple (y, f, i).
func (m YFI) Get() (yards, feet, inches int) {
	yards = int(m) / 36
	feet = (int(m) - yards * 36) / 12
	inches = (int(m) - yards * 36 - feet * 12)
	return
}

// Return a YFI value representing a distance of y yards, f feet, and i inches. 
func (m *YFI) Set(inches, feet, yards int) {
	*m = YFI(inches + feet * 12 + yards * 36)
}
