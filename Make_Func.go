/*
Function Declartion
		func function_name( [parameter list] ) [return_types]
		{
		   body of the function
		}
*/
package main
import "fmt"

func main() {
   /* local variable definition */
   var a int = 100
   var b int = 200
   var ret int
   /* calling a function to get max value */
   ret = max(a, b)
	 fmt.Println("\nMax Function")
   fmt.Printf( "Max value is : %d\n", ret )
	 fmt.Println("\n\nSwaping Function\n")

	 var m,n string = "Dishant","Raut";
	 fmt.Println("m = ", m)
	 fmt.Println("n = ", n)

	 m,n = swap(m,n)
	 fmt.Println("After Swaping")
	 fmt.Println("m = ", m)
	 fmt.Println("n = ", n)
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
   /* local variable declaration */
   var result int
   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}

func swap(x, y string) (string, string) {
   return y, x
}
