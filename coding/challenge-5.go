package coding

import (
	"strconv"
	"strings"
)

func InterpretArithmeticCommand(commands []string) (output float64) {
	// Code start here

	for _, command := range commands {
		part := strings.Split(command, " ")
		if len(part) != 2 {
			continue
		}

		op := part[0]
		num, err := strconv.Atoi(part[1])
		if err != nil {
			continue
		}

		newNum := float64(num)

		switch op {
		case "add":
			output += newNum

		case "multiply":
			output *= newNum

		case "substract":
			output -= newNum

		case "devide":
			if newNum != 0 {
				output /= newNum
			}
		}
	}

	return output
}
