
package main

import("fmt";"time")

func main() {

  t := time.Now()

  var no_day int
  fmt.Print("Enter Number = ")
  fmt.Scan(&no_day)

  switch no_day {
  case 1:
    fmt.Println("\nSunday")
  case 2:
    fmt.Println("\nMonday")
  case 3:
    fmt.Println("\nTuesday")
  case 4:
    fmt.Println("\nWednesday")
  case 5:
    fmt.Println("\nThursday")
  case 6:
    fmt.Println("\nFriday")
  case 7:
    fmt.Println("\nSaturday")
  }

  // To get todays timings
  fmt.Print("\n\nToday = ",t,"\n")

}
