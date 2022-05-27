package main

import "fmt"

func main(){
	var a,b float32

	fmt.Println("Birinchi  sondi kiriting")
	fmt.Scanf("%f",&a)

	fmt.Println("Ikkinnchi sondi kiriting")
	fmt.Scanf("%f",&b)

	if  a>b {
		fmt.Println(a*(-1),b)
	}else {
		fmt.Println(a,b)

	}


}