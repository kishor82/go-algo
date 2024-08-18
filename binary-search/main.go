package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	result := binarySearch(array, 9)
	fmt.Println("data", result)
}

func binarySearch(array []int, needle int) bool {
	lo, hi := 0, len(array)

	for lo < hi {
		middle := lo + (hi-lo)/2
		value := array[middle]

		switch {
		case value == needle:
			return true
		case value > needle:
			hi = middle
		default:
			lo = middle + 1
		}
	}

	return false
}
