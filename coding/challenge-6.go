package coding

func FindMaxSum(matrix [][]int) (maxSum int) {
	// Code start here

	newMatrix := len(matrix)

	// 5x5 sum
	for i := 0; i < newMatrix; i++ {
		for j := 0; j <= newMatrix-5; j++ {
			currentSum := 0
			for k := 0; k < 5; k++ {
				currentSum += matrix[i][j+k]
			}

			// bigger sum will be inputted in the max sum
			if currentSum > maxSum {
				maxSum = currentSum
			} else if currentSum <= maxSum {
				maxSum = currentSum
			}

		}

	}

	for i := 0; i <= newMatrix-5; i++ {
		for j := 0; j < newMatrix; j++ {
			currentSum := 0
			for k := 0; k < 5; k++ {
				currentSum += matrix[i+k][j]
			}
		}
	}

	return
}
