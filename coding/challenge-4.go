package main

func TrimSpaces(text string) (output string) {
	// Code start here

	output = ""

	inWord := false
	for i := 0; i < len(text); i++ {
		char := text[i]
		if char != ' ' {
			output += string(char)
			inWord = true
		} else if inWord {
			output += " "
			inWord = false
		}

	}
	if len(output) >= 0 && output[len(output)-1] == ' ' {
		output = output[:len(output)-1]
	}

	return output

}
