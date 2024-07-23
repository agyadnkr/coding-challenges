package coding

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
	t.Log("FIRST CASE: Simple String with 1 Vowel")

	text = "a"
	expectedResult = []int{1}

	result = FindVowelPosition(text)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End First Case

	// Second Case
	t.Log("SECOND CASE: Simple String with no Vowel")

	text = "bswfkplhv"
	expectedResult = []int{}

	result = FindVowelPosition(text)

	if len(result) > 0 {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Second Case

	// Third Case
	t.Log("THIRD CASE: Vowels with Whitespace")

	text = "Hello World!"
	expectedResult = []int{2, 5, 8}

	result = FindVowelPosition(text)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Third Case

	// Fourth Case
	t.Log("FOURTH CASE: Mixed Case Vowels")

	text = "HEllo WOrLd!"
	expectedResult = []int{2, 5, 8}

	result = FindVowelPosition(text)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Fourth Case

	// Fifth Case
	t.Log("FIFTH CASE: Text with Numbers and Symbols")

	text = "Th3r3 @r3 v0w3ls h3r3!"
	expectedResult = []int{}

	result = FindVowelPosition(text)

	if len(result) > 0 {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Fifth Case

	// Sixth Case
	t.Log("SIXTH CASE: Long Text")

	text = "This is a very long string with multiple vowels spread throughout the text."
	expectedResult = []int{3, 6, 9, 12, 17, 24, 29, 34, 37, 40, 43, 45, 52, 53, 59, 60, 63, 64, 69, 72}

	result = FindVowelPosition(text)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("INCCORECT RESULT: got: %v, want: %v.", result, expectedResult)
	}
	// End Sixth Case
}
