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

	if a<b && b<c {
		a=a*2
		b=b*2
		c=c*2
		fmt.Println(a,b,c)
	} else {
		a=a*(-1)
		b=b*(-1)
		c=c*(-1)
		fmt.Println(a,b,c)
	}
}