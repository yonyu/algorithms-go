package main

import (
	"fmt"
	"github.com/yonyu/algorithms-go/adt"
)

func main() {
	fmt.Println("Hello, we are going to write some algorithms in Go!")

	x := adt.Integer(5)
	var y adt.Integer = 10

	fmt.Println(x.IsOdd())
	fmt.Println(y.IsOdd())

	yfi1 := adt.YFI(100)
	fmt.Println(yfi1.Get())

	yfi1.Set(10, 10, 10)
	fmt.Println(yfi1.Get())

	adt.DemonstrateOverride()

	adt.TestNewBulkStore()

	adt.TestNewGroceryStoreBin()
}
