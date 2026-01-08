package main

func reduce(slice []int, initializer int, op func(int, int) int) int {
	acc := initializer
	for _, v := range slice {
		acc = op(acc, v)
	}
	return acc
}
