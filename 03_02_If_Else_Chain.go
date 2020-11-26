
package main

import("fmt")

func main() {
  
  fmt.Println("Enter your percentage = ")
  var per float32
  fmt.Scanln(&per)

  if ( per < 1 || per > 99 ){
    fmt.Print("\nInvalid Input...!!")
  }else if ( per > 0 && per < 50){
    fmt.Print("\nResult = Failed")
  }else if ( per >= 50 && per < 60){
    fmt.Print("\nResult = D grade")
  }else if ( per >= 60 && per < 70){
    fmt.Print("\nResult = C grade")
  }else if ( per >= 70 && per < 80){
    fmt.Print("\nResult = B grade")
  }else if ( per >= 80 && per < 90){
    fmt.Print("\nResult = A grade")
  }else{
    fmt.Print("\nResult = A+ grade")
  }
  
}
