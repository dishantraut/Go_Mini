package main

import "fmt"

func main() {
  // You can even ignore the length of the array in the declaration and replace it with ...
  // and let the compiler find the length for you
  a1 := [...]string{"Dishant", "23", "Data Analyst", "ML & AI", "Chat Bot Development", "Python & Go Lang"}
  for i := 1 ; i < len(a1) ; i++ {
    fmt.Println(i,". ",a1[i])
  }
}
