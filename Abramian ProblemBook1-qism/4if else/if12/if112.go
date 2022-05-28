package main

import "fmt"

func main(){
	var a,b,c float64

	fmt.Println("Birinchi sondi kiriting")
	fmt.Scanf("%f",&a)
	fmt.Println("Ikkinchi sondi kiriting")
	fmt.Scanf("%f",&b)
	fmt.Println("Uchinchi sondi kiriting")
	fmt.Scanf("%f",&c)

	if a>b && b>c {
		fmt.Println("Eng kichik son:",c)
	} else if c>a && a>b{
		fmt.Println("Eng kichik son:",b)
	} else {
		fmt.Println("Eng kichik son:",a)
	}
}