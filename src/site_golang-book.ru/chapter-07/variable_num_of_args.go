package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5}
	fmt.Println(add(xs...))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
