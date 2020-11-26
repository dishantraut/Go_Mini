
package main

import("fmt")

func main() {

  var user_id string
  var pwd string

  fmt.Print("Enter User ID = ")
  fmt.Scan(&user_id)
  fmt.Print("Enter Password = ")
  fmt.Scan(&pwd)

  if ( user_id == "Dishant1997" && pwd == "d!sh@nt97"){
    fmt.Println("Access Granted..!!")
  }else{
    fmt.Println("Access Denied..!!")
  }
}
