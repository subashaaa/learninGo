package main

import (
  "fmt"
)

type Vertex struct {
  Lat, Long float64
}

// var m map[string]Vertex

func main() {
  // m = make(map[string]Vertex)
  // m["Bell Labs"] = Vertex{
  //   40.68433, -74.39967,
  // }
  // var m = map[string]Vertex{
  //   "Bell Labs": Vertex{
  //     40.68433, -74-39967,
  //   },
  //   "Google": Vertex{
  //     37.42202, -122.08408,
  //   },
  // }
  var m = map[string]Vertex{
    "Bell Labs": {
      40.68433, -74-39967,
    },
    "Google": {
      37.42202, -122.08408,
    },
  }
  fmt.Println(m["Bell Labs"])
  fmt.Println(m)

  mm := make(map[string]int)

  mm["Answer"] = 42
  fmt.Println("The value:", mm["Answer"])

  mm["Answer"] = 48
  fmt.Println("The value:", mm["Answer"])
  
  delete(mm, "Answer")
  fmt.Println("The value:", mm["Answer"])

  v, ok := mm["Answer"]
  fmt.Println("The value:", v, "Present?", ok)
}
