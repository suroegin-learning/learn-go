package main

import "fmt"

type Person struct {
  Name string
}

func (p *Person) Talk() {
  fmt.Println("He, my name is", p.Name)
}

type Android struct {
  Person
  Model string
}

func main() {
  p := Person{"Ivan"}
  a := Android{p, "lol"}
  a.Person.Talk()
}
