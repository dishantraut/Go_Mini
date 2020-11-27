package main

import (
	"fmt"
)


func main() {
	// For & If
	fmt.Println("\n")
	for i := 1 ; i <= 10 ; i++ {
		if i % 2 == 0{
			fmt.Println("Even No = ",i)
		}else{
			fmt.Println("Odd No = ",i)
		}
	}
}
