package main

import (
  "fmt"
)

func fact(n int) int {
  if (n == 1) {
    return 1
  }else{
    return n*fact(n-1)
  }
}

func main() {
  var num int
  fmt.Print("Enter a number = ",)
  fmt.Scan(&num)
  fmt.Println("\nFactorial of ",num," = ",fact(num))
}
