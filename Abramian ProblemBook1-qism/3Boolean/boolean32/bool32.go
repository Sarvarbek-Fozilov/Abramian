package main

import "fmt"

func main(){
	var a,b,c int
	fmt.Println("Birinchi sondi kiriting")
	fmt.Scanf("%d",&a)
	fmt.Println("Ikkinchi sondi kiriting")
	fmt.Scanf("%d",&b)
	fmt.Println("Uchinchi sondi kiriting")
	fmt.Scanf("%d",&c)

	if (a*a==b*b+c*c) || (b*b==a*a+c*c) || (c*c==a*a+b*b){
		fmt.Println("To'g'ri burchakli uchburchak")
	} else{
		fmt.Println("boshqa qiymat kiritib ko'ring")
	}


}