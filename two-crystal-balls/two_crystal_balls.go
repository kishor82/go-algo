package main

import "math"

// Time complexity: O(Sqrt(N))

func twoCrystalBalls(breaks []bool) int {
	// jump Sqrt(N)
	jumpTo := math.Sqrt(float64((len(breaks))))
	i := int(jumpTo)

	for ; i < len(breaks); i += int(jumpTo) {
		if breaks[int(i)] {
			break
		}
	}

	// if crystal breaks try again from the previous steps so just moveback to (jumpTo) times
	i -= int(jumpTo)

	// We are scanning Sqrt(N) items
	for ; i < len(breaks); i++ {
		if breaks[i] {
			return i
		}
	}

	return -1
}
