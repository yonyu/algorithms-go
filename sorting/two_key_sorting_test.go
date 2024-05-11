package sorting

import (
	"fmt"
	"testing"
)

func TestTwoKeySorting(t *testing.T) {
	people := []Person {
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
		{"Nancy", "Bobdaughter", 19},
	}

	people = TwoKeySorting(people)
	fmt.Println(people)
}

