package main

import (
	"fmt"
	"main/coding"
)

func main() {

	fmt.Println(coding.FilterNameByValueThreshold([]string{"jason", "kimmy", "aiden"}, []int{1, 5, 10}, 3))

	fmt.Println(coding.FindVowelPosition("hello world"))

	// fmt.Println(coding.FindRepeatedNumber([1,2,3,3,4,5]))
}
