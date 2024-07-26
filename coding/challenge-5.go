package coding

import (
	"strconv"
	"strings"
)

func InterpretArithmeticCommand(commands []string) (output float64) {
	// Code start here

	output = 0

	for _, command := range commands {
		part := strings.Split(command, " ")
		if len(part) != 2 {
			continue
		}

		op := part[0]
		num, err := strconv.ParseFloat(part[1], 64)
		if err != nil {
			continue
		}

		switch op {
		case "Add":
			output += num

		case "Multiply":
			output *= num

		case "Subtract":
			output -= num

		case "Divide":
			if num != 0 {
				output /= num
			}
		}
	}

	return output
}
