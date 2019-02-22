package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(binarySearch(5, []int{6, 9, 1, 5, 7, 3, 2, 90}))
	fmt.Println(binarySearch(3, []int{1, 3, 5}))
}

func binarySearch(target int, collection []int) int {
	sort.Ints(collection)
	lowerbound := 0
	upperbound := len(collection)

	for {
		if upperbound < lowerbound {
			return -1
		}

		midValue := lowerbound + ((upperbound - lowerbound) / 2)

		switch {
		case collection[midValue] < target:
			lowerbound = midValue + 1
		case collection[midValue] > target:
			upperbound = midValue - 1
		case collection[midValue] == target:
			return midValue
		default:
			return -1

		}
	}
}
