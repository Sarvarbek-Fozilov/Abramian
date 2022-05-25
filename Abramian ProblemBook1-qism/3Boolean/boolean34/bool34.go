package main

import "fmt"

func main(){
	var x,y int

	fmt.Println("Shaxmat doskasining x kordintasi qiymatni kirit")
	fmt.Scanf("%d",&x)
	fmt.Println("Shaxmat doskasining y kordintasi qiymatni kirit")
	fmt.Scanf("%d",&y)
	

	if (x+y)%2!=0 {
		fmt.Printf("Tog'ri\n")
	} else {
		fmt.Printf("notog'ri\n")
	}
}