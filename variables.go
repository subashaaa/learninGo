package main

import (
  "fmt"
  "math/rand"
  "math"
  "time"
  "math/cmplx"
)

func add(x int, y int) int {
  return x + y
}

func swap(x, y string) (string, string) {
  return y, x
}

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

var c, python, java bool
var d, scala, ruby int = 1, 3, 5

var (
  ToBe bool = false
  MaxInt uint64 = 1 << 64 - 1
  z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Println("Pi:", math.Pi)
  fmt.Println("42 + 13 =", add(42, 13))
  
  a, b := swap("hello", "world")
  fmt.Println(a, b)
  
  fmt.Println(split(28))

  var i int
  fmt.Println(i, c, python, java)
  fmt.Println(i, d, scala, ruby)

  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)

  j := 42           // var j int = 42
  f := float64(j)   // var f float64 = float64(j)
  u := uint(f)      // var u uint = uint(f)
  fmt.Printf("%T %v %T %v %T %v\n", j, j, f, f, u, u)
}   
