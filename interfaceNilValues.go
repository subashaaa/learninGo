package main

import (
  "fmt"
)

type I interface {
  M()
}

func describe(i I) {
  fmt.Println("(%v, %T)\n", i, i)
}

func main() {
  var i I
  describe(i)
  i.M()
}