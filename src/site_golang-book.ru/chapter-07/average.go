package main

import "fmt"

func main() {
	xs := []float64{22, 51,76, 12, 91}

	total := 0.0
	for _, value := range xs {
		total += value
	}
	fmt.Println(total / float64(len(xs)))
}

