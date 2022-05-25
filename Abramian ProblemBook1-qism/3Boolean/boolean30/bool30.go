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

	if a==b && b==c {
		fmt.Println("Teng tomonli uchburchak")
	} else{
		fmt.Println("boshqa qiymat kiritib ko'ring")
	}


}