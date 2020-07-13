package main

import "fmt"

func main() {
	x:=0
   for true  {
       fmt.Printf("This loop will run forever.\n");
			 if(x==3){
				 break
			 }else{
				fmt.Println("Yoo",x)
			 }
			 x+=1
	 }
}
