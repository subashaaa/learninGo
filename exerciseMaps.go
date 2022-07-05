package main

import (
	// "golang.org/x/tour/wc"
	"strings"
  "fmt"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words{
		m[v] += 1
	}
	
	return m
}

func main() {
	// wc.Test(WordCount)
  fmt.Println(WordCount("I am learning Go!"))
  fmt.Println(WordCount("The quick brown fox jumped over the lazy dog."))
  fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
  fmt.Println(WordCount("A man a plan a canal panama."))
}
