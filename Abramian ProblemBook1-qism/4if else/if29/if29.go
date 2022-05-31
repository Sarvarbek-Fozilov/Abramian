package main

import "fmt"

func main(){
	var num int

	fmt.Println("Son kiritng")
	fmt.Scanf("%i",&num)

		 if num>0 {
			fmt.Println("Musbat son")
		} else if num<0 {
			fmt.Println("manfiy ")
		} else if num%2==0 {
			fmt.Println("Juft")
		} else {
			fmt.Printf("Toq son")
		}
	

}