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
		fmt.Println("O'rtacha son:",b)
	} else if c>a && a>b{
		fmt.Println("O'rtacha son:",a)
	} else {
		fmt.Println("O'rtacha son:",c)
	}
}