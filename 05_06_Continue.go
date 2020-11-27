package main

import (
	"fmt"
)


func main() {
	// For & If
	fmt.Println("\n")
	for i := 1 ; i <= 10 ; i++ {
		if i % 2 == 0{
			continue
		}else{
			fmt.Println("No = ",i)
		}
	}
}
