
package main

import ("fmt")

// Variable Declerations & Assingments
func main() {
  // Basic Syntax For Deacration
  /* var identifier type */

    // Assingments & Declerations together
  var age1 int = 23
  var name1 string = "Dishant"

    // Initialization
  var age2 int
  var name2 string
  age2 = 48
  name2 = "Rajendra"

    // Short Declerations
    // https://www.geeksforgeeks.org/difference-between-var-keyword-and-short-declaration-operator-in-golang/
  name3 := "Priyanshu"
  age3 := 20

  fmt.Println("\nName =  ",name1)
  fmt.Println("Age =  ",age1)
  fmt.Println("\nName = ",name2)
  fmt.Println("Age = ",age2)
  fmt.Println("\nName = ",name3)
  fmt.Println("Age = ",age3)

}
