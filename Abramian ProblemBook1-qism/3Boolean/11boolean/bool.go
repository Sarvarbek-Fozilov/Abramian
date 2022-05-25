package main

import "fmt"

func main(){
	  var  a,b int
	  fmt.Println("A sonini kirting ")
	  fmt.Scanf("%d",&a)
	  fmt.Println("B sonini kirting")
	  fmt.Scanf("%d",&b)
	  if a%2!=0 && b%2!=0{
		  fmt.Println("ikkisi ham toq son")
	  } else if  a%2==0 && b%2==0 {
		  fmt.Println("Ikkisi ham juft  son")
	  } else {
		  fmt.Println("Shart qanoatlantrmadi.")
	  }


	
}