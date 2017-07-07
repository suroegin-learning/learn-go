package main

import "fmt"

func swap(x, y *int) {
  *x, *y = *y, *x
}

func main() {
  x, y := 2, 3

  fmt.Println("x =", x, "y =", y)
  swap(&x, &y)
  fmt.Println("Now x =", x, "y =", y)
}
