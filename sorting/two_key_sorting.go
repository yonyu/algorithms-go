package sorting

import (
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func TwoKeySorting(people []Person) []Person {
	sort.Slice(people, func(i, j int ) bool {
    if people[i].LastName == people[j].LastName {
			return people[i].FirstName < people[j].FirstName
		} else {
			return people[i].LastName < people[j].LastName
		}
	})

	return people
}

func TwoKeyStableSorting(people []Person) []Person {
	sort.SliceStable(people, func(i, j int ) bool {
    if people[i].LastName == people[j].LastName {
			return people[i].FirstName < people[j].FirstName
		} else {
			return people[i].LastName < people[j].LastName
		}
	})
	
	return people
}