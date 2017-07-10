package main

import "fmt"
import "os"
import "strconv"

func main() {
  fmt.Println(fib(strconv.Atoi(os.Args[1:])))
}

func fib(n int) int {
  x, y := 0, 1
  for i:=0; i<n; i++ {
    x, y = y, x+y
  }
  return x
}
