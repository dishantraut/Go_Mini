
package main

import("fmt")

func main() {

  num1 := 20
  num2 := 20
  
  if( num1 > num2 ){
    fmt.Println("\nNum1 is greater = ",num1)
  }else if ( num1 < num2 ){
    fmt.Println("\nNum2 is greater = ",num2)
  }else{
    fmt.Println("\nBoth are Equal \nThank You..!!")
  }
}
