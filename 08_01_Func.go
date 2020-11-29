package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func hw(name string) {
  fmt.Println("\n\nHello ",name)
}
func main() {
  var num1 int
  var num2 int
  fmt.Println("Enter Num1 = ")
  fmt.Scan(&num1)
  fmt.Println("Enter Num2 = ")
  fmt.Scan(&num2)
  hw("Dishant")
	fmt.Println("Addition = ",add(num1, num2))
}
