package main

func binarySearch(input []int, needle int) bool {
	low, high := 0, len(input)

	for low < high {
		middle := low + (high-low)/2
		value := input[middle]

		switch {
		case value == needle:
			return true
		case value > needle:
			high = middle
		default:
			low = middle + 1
		}
	}

	return false
}
