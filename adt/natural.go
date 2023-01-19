package adt

type Integer int

type Natural Integer

func (n Integer) IsOdd() bool { // (n Integer): receiver
	return n % 2 == 1
}

func (n Natural) Succ() Natural {
	if n < 0 {
		panic("value is less than 0")
	}
	return n + 1
}

func (n Natural) Pred() Natural {
	if n < 1 {
		panic("value is less than 1")
	}

	return n - 1
}