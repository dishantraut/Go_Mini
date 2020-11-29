package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}
func main() {
  a := "Dishant"
  b := "Priyanshu"
  fmt.Println("a = ",a," b = ",b)
  a1 , b1 := swap(a,b)
  fmt.Println("a = ",a1," b = ",b1)
}
