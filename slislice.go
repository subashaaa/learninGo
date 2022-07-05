package main

import (
  "fmt"
  "strings"
)

func printSlice(x []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}

func main() {
  board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }

  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }

  var s []int
  printSlice(s)

  s = append(s, 0)
  printSlice(s)
  
  s = append(s, 1)
  printSlice(s)

  s = append(s, 2)
  printSlice(s)
  
  s = append(s, 3)
  printSlice(s)

  s = append(s, 4)
  printSlice(s)
}
