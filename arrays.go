package main

import (
  "fmt"
)


func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)

  var sliceOfInts []int = primes[1:4]
  fmt.Println(sliceOfInts)

  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
  }
  fmt.Println(names)

  c := names[0:2]
  d := names[1:3]
  fmt.Println(c, d)

  d[0] = "XXX"
  fmt.Println(c, d)
  fmt.Println(names)

  q := []int{2, 3, 5, 7, 11, 13}
  fmt.Println(q)

  r := []bool{true, false, true, true, false, true}
  fmt.Println(r)

  s := []struct{
    i int
    b bool
  }{
    {2, true},
    {3, false},
    {5, true},
    {7, true},
    {11, false},
    {13, true},
  }
  fmt.Println(s)
  
  w := q 

  w = w[1:4]
  fmt.Println(w)

  w = w[:2]
  fmt.Println(w)
  
  w = w[1:]
  fmt.Println(w)

  y := []int{2, 3, 5, 7, 11, 13}
  printSlice(y)

  y = y[:0]
  printSlice(y)

  y = y[:4]
  printSlice(y)

  y = y[2:]
  printSlice(y)
}
