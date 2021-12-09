package main

import "fmt"

func main() {
	var nums = make([]int, 11)
	var s string
	for i := range nums {
		nums[i] = i
		if i%2 == 0 {
			s = "even"
		} else {
			s = "odd"
		}
		fmt.Printf("%v is %s\n", nums[i], s)
	}
}
