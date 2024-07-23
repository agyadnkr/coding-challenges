package coding

func FindVowelPosition(text string) (output []int) {
	// Code start here

	// huruf vokal a, i, u, e, o

	var vowel = []string{"a", "i", "u", "e", "o"}

	var hasilArray []int

	for i := 0; i < len(text); i++ {
		for j := 0; j < len(vowel); j++ {
			if vowel[j] == string(text[i]) {
				hasilArray = append(hasilArray, i+1)
			}
		}
	}

	return hasilArray
}
