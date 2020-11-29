package main

import "fmt"

func main() {
  // You can even ignore the length of the array in the declaration and replace it with ...
  // and let the compiler find the length for you
  a1 := [...]int{12, 78, 50}
  f := [...]float64{67.7, 89.8, 21, 78}
  a := [...]string{"USA", "China", "India", "Germany", "France"}
  b := a // a copy of a is assigned to b
  b[0] = "Singapore"
  fmt.Println("a1 is ",a1)
  fmt.Println("a is  ", a)
  fmt.Println("b is  ", b)
  fmt.Println("Length of 'f' is",len(f))
}
