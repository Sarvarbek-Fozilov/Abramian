package main

import "fmt"

func main(){
	var a,b int

	fmt.Println("Birinchi  sondi kiriting")
	fmt.Scanf("%d",&a)

	fmt.Println("Ikkinnchi sondi kiriting")
	fmt.Scanf("%d",&b)

	if a!=b{
		a=a+b
		b=a+b
		fmt.Println(a,b)
	
	}else {
		a=0
		b=0
		fmt.Println(a,b)

	}


}