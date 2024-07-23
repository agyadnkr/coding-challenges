package main

import (
	"fmt"
	"main/coding"
)

func main() {

	fmt.Println(coding.FilterNameByValueThreshold([]string{"jason", "kimmy", "aiden"}, []int{1, 5, 10}, 3))

}
