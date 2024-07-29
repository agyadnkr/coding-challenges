package main

func FilterNameByValueThreshold(names []string, nameValues []int, threshold int) (output []string) {
	// Code start here

	var group []string

	for i, values := range nameValues {
		if values > threshold {
			group = append(group, names[i])
		}

	}

	return group

}
