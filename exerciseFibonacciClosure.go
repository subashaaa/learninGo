package main

import "fmt"

func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		// temp := first
		// first = second
		// second = temp + second
		first, second = second, first + second
    return second - first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
