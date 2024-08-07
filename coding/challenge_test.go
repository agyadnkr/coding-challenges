package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilterNameByValueThreshold(t *testing.T) {
	var (
		names          []string
		nameValues     []int
		threshold      int
		expectedResult []string
		result         []string
	)

	// First Case
	t.Log("FIRST CASE: Single Pair")

	names = []string{"Alice"}
	nameValues = []int{50}
	threshold = 40
	expectedResult = []string{"Alice"}

	result = FilterNameByValueThreshold(names, nameValues, threshold)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %+q, want: %+q.", result, expectedResult)
	}
	// End First Case

	// Second Case
	t.Log("SECOND CASE: No Matches")

	names = []string{"Alice", "Ilham", "Paul"}
	nameValues = []int{50, 20, 12}
	threshold = 51
	expectedResult = []string{}

	result = FilterNameByValueThreshold(names, nameValues, threshold)

	if len(result) > 0 {
		t.Errorf("INCCORECT RESULT: got: %+q, want: %+q.", result, expectedResult)
	}
	// End Second Case

	// Third Case
	t.Log("THIRD CASE: All Matches")

	names = []string{"Alice", "Ilham", "Paul", "Zorro", "Luffy", "Sanji"}
	nameValues = []int{50, 20, 12, 22, 43, 10}
	threshold = 5
	expectedResult = []string{"Alice", "Ilham", "Paul", "Zorro", "Luffy", "Sanji"}

	result = FilterNameByValueThreshold(names, nameValues, threshold)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %+q, want: %+q.", result, expectedResult)
	}
	// End Third Case

	// Fourth Case
	t.Log("FOURTH CASE: Mixed Matches")

	names = []string{"Alice", "Ilham", "Paul", "Zorro", "Luffy", "Sanji"}
	nameValues = []int{50, 20, 12, 22, 43, 10}
	threshold = 21
	expectedResult = []string{"Alice", "Zorro", "Luffy"}

	result = FilterNameByValueThreshold(names, nameValues, threshold)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %+q, want: %+q.", result, expectedResult)
	}
	// End Fourth Case

	// Fifth Case
	t.Log("FIFTH CASE: Values Equal to Threshold")

	names = []string{"Alice", "Ilham", "Paul", "Zorro", "Luffy", "Sanji"}
	nameValues = []int{50, 50, 50, 50, 50, 50}
	threshold = 50
	expectedResult = []string{}

	result = FilterNameByValueThreshold(names, nameValues, threshold)

	if len(result) > 0 {
		t.Errorf("INCCORECT RESULT: got: %+q, want: %+q.", result, expectedResult)
	}
	// End Fifth Case

	// Sixth Case
	t.Log("SIXTH CASE: Negative Values")

	names = []string{"Alice", "Ilham", "Paul", "Zorro", "Luffy", "Sanji"}
	nameValues = []int{-25, -12, -22, -10, -8, -55}
	threshold = -20
	expectedResult = []string{"Ilham", "Zorro", "Luffy"}

	result = FilterNameByValueThreshold(names, nameValues, threshold)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %+q, want: %+q.", result, expectedResult)
	}
	// End Sixth Case

	// Seventh Case
	t.Log("SEVENTH CASE: Large Input Size")

	threshold = 50000
	names = []string{}
	nameValues = []int{}
	expectedResult = []string{}

	for i := 1; i <= 1000000; i++ {
		names = append(names, fmt.Sprintf("Name %d", i))
		nameValues = append(nameValues, i)

		if i > threshold {
			expectedResult = append(expectedResult, fmt.Sprintf("Name %d", i))
		}
	}

	result = FilterNameByValueThreshold(names, nameValues, threshold)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: result length: %d, expected result length: %d.", len(result), len(expectedResult))
	}
	// End Seventh Case
}

func TestFindVowelPosition(t *testing.T) {
	var (
		text           string
		result         []int
		expectedResult []int
	)

	// First Case
	text = "a"
	expectedResult = []int{1}

	t.Logf("FIRST CASE: Simple String with 1 Vowel (%s)\n", text)

	result = FindVowelPosition(text)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End First Case

	// Second Case
	text = "bswfkplhv"
	expectedResult = []int{}

	t.Logf("SECOND CASE: Simple String with no Vowel (%s)\n", text)

	result = FindVowelPosition(text)

	if len(result) > 0 {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Second Case

	// Third Case
	text = "Hello World!"
	expectedResult = []int{2, 5, 8}

	t.Logf("THIRD CASE: Vowels with Whitespace (%s)\n", text)

	result = FindVowelPosition(text)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Third Case

	// Fourth Case
	text = "HEllo WOrLd!"
	expectedResult = []int{2, 5, 8}

	t.Logf("FOURTH CASE: Mixed Case Vowels (%s)\n", text)

	result = FindVowelPosition(text)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Fourth Case

	// Fifth Case
	text = "Th3r3 @r3 v0w3ls h3r3!"
	expectedResult = []int{}

	t.Logf("FIFTH CASE: Text with Numbers and Symbols (%s)\n", text)

	result = FindVowelPosition(text)

	if len(result) > 0 {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Fifth Case

	// Sixth Case
	text = "This is a very long string with multiple vowels spread throughout the text."
	expectedResult = []int{3, 6, 9, 12, 17, 24, 29, 34, 37, 40, 43, 45, 52, 53, 59, 60, 63, 64, 69, 72}
	t.Logf("SIXTH CASE: Long Text (%s)\n", text)

	result = FindVowelPosition(text)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Sixth Case
}

func TestFindRepeatedNumber(t *testing.T) {
	var (
		arr            []int
		n              int
		result         []int
		expectedResult []int
	)

	// First Case
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	n = 1
	expectedResult = []int{1, 2, 3, 4, 5, 6, 7, 8}

	t.Logf("FIRST CASE: Single Occurence (n: %d arr: %v)\n", n, arr)

	result = FindRepeatedNumber(arr, n)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End First Case

	// Second Case
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	n = 2
	expectedResult = []int{}

	t.Logf("SECOND CASE: No Occurence (n: %d arr: %v)\n", n, arr)

	result = FindRepeatedNumber(arr, n)

	if len(result) > 0 {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Second Case

	// Third Case
	arr = []int{1, 2, 3, 4, 2, 3, 4, 5, 6, 4, 3, 2}
	n = 3
	expectedResult = []int{2, 3, 4}

	t.Logf("THIRD CASE: Multiple Matching Occurences (n: %d arr: %v)\n", n, arr)

	result = FindRepeatedNumber(arr, n)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Third Case

	// Fourth Case
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	n = 1
	expectedResult = []int{15}

	t.Logf("FOURTH CASE: Exact Occurences (n: %d arr: %v)\n", n, arr)

	result = FindRepeatedNumber(arr, n)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Fourth Case

	// Fifth Case
	arr = []int{}
	n = 7
	expectedResult = []int{}

	for i := 1; i <= 1000000; i++ {
		for j := 0; j < 5; j++ {
			arr = append(arr, i)
		}
	}

	t.Logf("FIFTH CASE: Large Array (n: %d arr: 1 - 1 million)\n", n)

	result = FindRepeatedNumber(arr, n)

	if len(result) > 0 {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Fifth Case
}

func TestTrimSpace(t *testing.T) {
	var (
		text           string
		result         string
		expectedResult string
	)

	// First Case
	text = "Hello World!"
	expectedResult = "Hello World!"

	t.Logf("FIRST CASE: Simple Text with No Double Spaces (%s)\n", text)

	result = TrimSpaces(text)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%s\", want: \"%s\".", result, expectedResult)
	}
	// End First Case

	// Second Case
	text = " Hello World! "
	expectedResult = "Hello World!"

	t.Logf("SECOND CASE: Text with Leading and Trailing Spaces (%s)\n", text)

	result = TrimSpaces(text)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%s\", want: \"%s\".", result, expectedResult)
	}
	// End Second Case

	// Third Case
	text = "       Hello        World!          "
	expectedResult = "Hello World!"

	t.Logf("THIRD CASE: Text with Multiple Whitespaces (%s)\n", text)

	result = TrimSpaces(text)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%s\", want: \"%s\".", result, expectedResult)
	}
	// End Third Case

	// Fourth Case
	text = "       Lorem ipsum dolor sit  amet,  consectetur  adipiscing elit. Proin facilisis ullamcorper     viverra. Maecenas vel venenatis dolor, sed vulputate mi. Donec vel        semper velit. Nam     mauris metus,     condimentum dapibus metus ut,     ornare    porttitor    enim.          "
	expectedResult = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin facilisis ullamcorper viverra. Maecenas vel venenatis dolor, sed vulputate mi. Donec vel semper velit. Nam mauris metus, condimentum dapibus metus ut, ornare porttitor enim."

	t.Logf("FOURTH CASE: Long Text with Multiple Whitespaces (%s)\n", text)

	result = TrimSpaces(text)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%s\", want: \"%s\".", result, expectedResult)
	}
	// End Fourth Case

}

func TestInterpretArithmeticCommand(t *testing.T) {
	var (
		commands       []string
		result         float64
		expectedResult float64
	)

	// First Case
	commands = []string{}
	expectedResult = 0

	t.Logf("FIRST CASE: Empty Command (%v)\n", commands)

	result = InterpretArithmeticCommand(commands)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%.2f\", want: \"%.2f\".", result, expectedResult)
	}
	// End First Case

	// Second Case
	commands = []string{"Add 5"}
	expectedResult = 5

	t.Logf("FIRST CASE: Single Command (%v)\n", commands)

	result = InterpretArithmeticCommand(commands)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%.2f\", want: \"%.2f\".", result, expectedResult)
	}
	// End Second Case

	// Third Case
	commands = []string{"Add 1", "Add 1", "Add 1", "Add 1", "Multiply 2", "Divide 2"}
	expectedResult = 4

	t.Logf("THIRD CASE: Simple Multiple Arithmetic (%v)\n", commands)

	result = InterpretArithmeticCommand(commands)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%.2f\", want: \"%.2f\".", result, expectedResult)
	}
	// End Third Case

	// Fourth Case
	commands = []string{"Add 10", "Subtract 5", "Subtract 5", "Divide 10"}
	expectedResult = 0

	t.Logf("FOURTH CASE: Zero Divided by Number (%v)\n", commands)

	result = InterpretArithmeticCommand(commands)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%.2f\", want: \"%.2f\".", result, expectedResult)
	}
	// End Fourth Case

	// Fifth Case
	commands = []string{"Add 10", "Subtract 5", "Subtract 5", "Subtract 2", "Divide 0"}
	expectedResult = -2

	t.Logf("FIFTH CASE: Division by Zero (%v)\n", commands)

	result = InterpretArithmeticCommand(commands)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%.2f\", want: \"%.2f\".", result, expectedResult)
	}
	// End Fifth Case

	// Sixth Case
	commands = []string{"add 10", "Subtract 5", "subtract 5", "Subtract 2", "add 10", "Add 20", "Multiply 8"}
	expectedResult = 104

	t.Logf("SIXTH CASE: Case-Sensitive Commands (%v)\n", commands)

	result = InterpretArithmeticCommand(commands)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%.2f\", want: \"%.2f\".", result, expectedResult)
	}
	// End Sixth Case

	// Seventh Case
	commands = []string{"Add 1", "Add 9", "Divide 2", "Multiply 10", "Divide two"}
	expectedResult = 50

	t.Logf("SEVENTH CASE: Command with Invalid Number (%v)\n", commands)

	result = InterpretArithmeticCommand(commands)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%.2f\", want: \"%.2f\".", result, expectedResult)
	}
	// End Seventh Case

	// Eighth Case
	commands = []string{"Add 10000", "Add 15000", "Divide 500", "Multiply 10000"}
	expectedResult = 500000

	t.Logf("EIGHTH CASE: Command with Large Number (%v)\n", commands)

	result = InterpretArithmeticCommand(commands)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%.2f\", want: \"%.2f\".", result, expectedResult)
	}
	// End Eighth Case

	// Ninth Case
	commands = []string{"Add 15", "Divide 2"}
	expectedResult = 7.5

	t.Logf("NINTH CASE: Decimal Result (%v)\n", commands)

	result = InterpretArithmeticCommand(commands)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%.2f\", want: \"%.2f\".", result, expectedResult)
	}
	// End Ninth Case
}

func TestFindMaxSum(t *testing.T) {
	var (
		matrix         [][]int
		result         int
		expectedResult int
	)

	// First Case
	matrix = [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 3, 1, 1},
		{1, 1, 4, 1, 1},
		{1, 1, 5, 1, 1},
	}
	expectedResult = 15

	t.Log("FIRST CASE: Simple Matrix with Vertical Max Sum")
	for _, row := range matrix {
		t.Log(row)
	}

	result = FindMaxSum(matrix)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%d\", want: \"%d\".", result, expectedResult)
	}
	// End First Case

	// Second Case
	matrix = [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 2, 3, 4, 5},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}
	expectedResult = 15

	t.Log("SECOND CASE: Simple Matrix with Horizontal Max Sum")
	for _, row := range matrix {
		t.Log(row)
	}

	result = FindMaxSum(matrix)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%d\", want: \"%d\".", result, expectedResult)
	}
	// End Second Case

	// Third Case
	matrix = [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36},
	}
	expectedResult = 170

	t.Log("THIRD CASE: Larger Matrix with Mixed Values")
	for _, row := range matrix {
		t.Log(row)
	}

	result = FindMaxSum(matrix)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%d\", want: \"%d\".", result, expectedResult)
	}
	// End Third Case

	// Fourth Case
	matrix = [][]int{
		{-1, -2, -3, -4, -5},
		{-6, -7, -8, -9, -10},
		{-11, -12, -13, -14, -15},
		{-16, -17, -18, -19, -20},
		{-21, -22, -23, -24, -25},
	}
	expectedResult = -15

	t.Log("FOURTH CASE: Matrix with Negative Values")
	for _, row := range matrix {
		t.Log(row)
	}

	result = FindMaxSum(matrix)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%d\", want: \"%d\".", result, expectedResult)
	}
	// End Fourth Case

	// Fifth Case
	matrix = [][]int{
		{1, -2, 3, -4, 5, -6, 7},
		{-8, 9, -10, 11, -12, 13, -14},
		{-17, 20, -23, 26, -29, 32, -35},
		{-26, 31, -36, 41, -46, 51, -56},
		{-35, 42, -49, 56, -63, -70, -77},
		{-44, 53, -62, 71, -80, 89, -98},
		{-53, 64, -75, 86, -97, 109, -119},
	}
	expectedResult = 280

	t.Log("FIFTH CASE: Matrix with Mixed Negative and Positive Values")
	for _, row := range matrix {
		t.Log(row)
	}

	result = FindMaxSum(matrix)

	if expectedResult != result {
		t.Errorf("INCCORECT RESULT: got: \"%d\", want: \"%d\".", result, expectedResult)
	}
	// End Fifth Case
}
