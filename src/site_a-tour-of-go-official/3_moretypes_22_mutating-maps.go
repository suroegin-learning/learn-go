package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["Answer"] = 42
  fmt.Println("The value:", m["Answer"], "\n", m)

  delete(m, "Answer")
  fmt.Println(m)

  v, ok := m["Answer"]
  fmt.Println("The value:", v, "Present?", ok)
}

