package sorting

// Bubble sort algorithm
func BubbleSort(input []int) []int {
	lastIndex := len(input) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < lastIndex; i++ {
			if input[i] > input[i+1] {
				sorted = false
				input[i], input[i+1] = input[i+1], input[i]
			}
		}

		lastIndex -= 1
	}

	return input
}
