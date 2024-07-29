package main

func FindRepeatedNumber(numberArr []int, n int) (output []int) {
	// Code start here

	counts := make(map[int]int)

	for _, num := range numberArr {
		counts[num]++
	}

	for i, num := range numberArr {
		if counts[num] == n {
			output = append(output, i+1)
			counts[num] = 0
		}
	}

	return output
}
