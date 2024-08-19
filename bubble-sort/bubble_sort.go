package main

func BubbleSort(elemets []int) []int {
	for i := 0; i < len(elemets); i++ {
		for j := 0; j < len(elemets)-1-i; j++ {
			if elemets[j] > elemets[j+1] {
				temp := elemets[j]
				elemets[j] = elemets[j+1]
				elemets[j+1] = temp
			}
		}
	}

	return elemets
}
