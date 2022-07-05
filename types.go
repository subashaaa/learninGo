package main

import (
  "fmt"
)

type Vertex struct {
  X int
  Y int
}

var (
  v1 = Vertex{1, 2}
  v2 = Vertex{X: 1}
  v3 = Vertex{}
  w = &Vertex{1, 2} 
)

func main() {
  i, j := 42, 2701

  p := &i
  fmt.Println(*p)
  *p = 21
  fmt.Println(i)

  p = &j
  *p = *p / 37
  fmt.Println(j)

  fmt.Println(Vertex{1, 2})

  v := Vertex{3, 4}
  v.X = 5
  fmt.Println(v.X, v.Y)

  z := &v
  z.X = 6
  fmt.Println(v)

  fmt.Println(v1, w, v2, v3)
}
