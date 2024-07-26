package coding

import "fmt"

func FindMaxSum(matrix [][]int) (maxSum int) {
	// Code start here

	newMatrix := len(matrix)

	// minimum mulai dari yang paling kecil  biar dapat nilai negatif yang besar
	maxSum = -1 << 5
	fmt.Println(maxSum)

	// 5x5 sum horizontal
	for i := 0; i < newMatrix; i++ {
		for j := 0; j <= newMatrix-5; j++ {
			currentSum := 0
			for k := 0; k < 5; k++ {
				currentSum += matrix[i][j+k]
			}

			// bigger sum will be inputted in the max sum
			if currentSum > maxSum {
				maxSum = currentSum
				fmt.Println("max sum horizontal", maxSum)
			}
		}

	}

	// vertical
	for i := 0; i <= newMatrix-5; i++ {
		for j := 0; j < newMatrix; j++ {
			currentSum := 0
			for k := 0; k < 5; k++ {
				currentSum += matrix[i+k][j]
			}

			if currentSum > maxSum {
				maxSum = currentSum
				fmt.Println("max sum vertical:", maxSum)
			}
		}
	}

	return
}
