package main

import "sort"

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Lucas", 30},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
}
