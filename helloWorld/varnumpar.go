package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7,9,3,5,1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d\n", x)
	printEach(slice...)
}

func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		fmt.Printf("%d\n", v)
		if v < min {
			min = v
		}
	}
	return min
}

func printEach(list ...int) {
	for _, v := range list {
		fmt.Printf("%d\n", v)
	}
}