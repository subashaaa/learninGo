package main

import (
  "fmt"
  "strconv"
  "math"
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
}

func about(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "About")
}

type Shape interface {
  area() float64
}

type Circle struct {
  x, y, radius float64
}

type Rectangle struct {
  width, height float64
}

func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
  return r.width * r.height
}

func getArea(s Shape) float64 {
  return s.area()
}

func adder() func(x int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func (p Person) greet() string {
  return "Hello, my name is " + p.firstName + " and my age is " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
  p.age++
}

type Person struct {
  firstName string
  lastName string
  city string
  gender string
  age int
}

func main() {
  var name = "Bob"
  var age float32 = 37.5

  fmt.Println(name, age)
  fmt.Printf("%T\n", age)

  var fruitArr [2]string
  fruitArr[0] = "Apple"
  fruitArr[1] = "Orange"

  fmt.Println(fruitArr)

  fruitSlice := []string{"Apple", "Orange", "Grape"}

  fmt.Println(fruitSlice, len(fruitSlice), cap(fruitSlice))

  fruitSlice = append(fruitSlice, "Djes")
  fmt.Println(fruitSlice, len(fruitSlice), cap(fruitSlice))

  x := 5
  y := 10

  if x < y {
    fmt.Printf("%d is less than %d\n", x, y)
  }

  for i := 1; i <= 10; i++ {
    fmt.Println(i)
  }

  for i := 1; i <= 100; i++ {
    if i % 15 == 0 {
      fmt.Println("FizzBuzz")
    } else if i % 3 == 0 {
      fmt.Println("Fizz")
    } else if i % 5 == 0 {
      fmt.Println("Buzz")
    } else {
      fmt.Println(i)
    }
  }

  emails := make(map[string]string)
  emails2 := map[string]string{
    "Bob": "bob@gmail.com",
    "Sharon": "sharon@gmail.com",
  }
  
  emails["Bob"] = "bob@gmail.com"
  emails["Sharon"] = "sharon@gmail.com"
  emails["Mike"] = "mike@gmail.com"

  fmt.Println(emails)
  fmt.Println(len(emails))
  fmt.Println(emails["Bob"])

  delete(emails, "Bob")

  fmt.Println(emails)
  fmt.Println(emails2)

  ids := []int{33, 76, 54, 23, 11, 2}

  for i, v := range ids {
    fmt.Printf("%d - ID: %d\n", i, v)
  }

  emails3 := map[string]string{
    "Bob": "bob@gmail.com",
    "Sharon": "sharon@gmail.com",
  }

  for k, v := range emails3 {
    fmt.Printf("%s: %s\n", k, v)
  }

  for k := range emails3 {
    fmt.Println("Name: " + k)
  }

  a := 5
  b := &a

  fmt.Println(a, b)
  fmt.Printf("%T\n", b)
  fmt.Println(*b)

  *b = 10
  fmt.Println(a)

  sum := adder()
  for i := 0; i < 10; i++ {
    fmt.Println(sum(i))
  }

  person1 := Person{
    firstName: "Samantha",
    lastName: "Smith",
    city: "Boston",
    gender: "f",
    age: 25,
  }

  fmt.Println(person1)
  fmt.Println(person1.greet())
  person1.hasBirthday()
  fmt.Println(person1.greet())

  circle := Circle{x:0, y:0, radius:5}
  rectangle := Rectangle{width:10, height:5}

  fmt.Printf("Circle area: %f\n", getArea(circle))
  fmt.Printf("Rectangle area: %f\n", getArea(rectangle))

  http.HandleFunc("/", index)
  http.HandleFunc("/about", about)
  fmt.Println("Server starting...")
  http.ListenAndServe(":3000", nil)
}
