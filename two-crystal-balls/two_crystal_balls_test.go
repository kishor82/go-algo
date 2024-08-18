package main

import (
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	t.Run("random breaking point", func(t *testing.T) {
		idx := rand.Intn(10000)
		data := make([]bool, 10000)
		for i := idx; i < 10000; i++ {
			data[i] = true
		}
		result := twoCrystalBalls(data)
		if result != idx {
			t.Errorf("Expected %d, but got %d", idx, result)
		}
	})

	t.Run("no breaking point", func(t *testing.T) {
		data := make([]bool, 821)
		result := twoCrystalBalls(data)
		if result != -1 {
			t.Errorf("Expected -1, but got %d", result)
		}
	})
}
