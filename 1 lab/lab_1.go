package main

import (
	"fmt"
	"sort"
)

func main() {
	integers := []int{10, -1, 20, 0, -100, 15, -8, -7, 1, 2, 5, 3, 4}
	k := 5

	sKResult := smallestK(integers, k)
	fmt.Println(fmt.Sprintf("k smallest elem %v", sKResult))

	mKResult := largestK(integers, k)
	fmt.Println(fmt.Sprintf("k largest elem %v", mKResult))
}

func smallestK(arr []int, k int) []int {
	sort.Ints(arr)

	return arr[:k]
}

func largestK(arr []int, k int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	return arr[:k]
}
