package main

import (
	"fmt"


)

func main(){
    var a,b,c,d float64
	fmt.Println("Tenglamaning birinchi sonini kiriting")
	fmt.Scanf("%d",&a)
	fmt.Println("Tenglamaning ikkinchi sonini kiriting")
	fmt.Scanf("%d",&b)
	fmt.Println("Tenglamaning uchinchi sonini kiriting")
	fmt.Scanf("%d",&c)

	d= b*b -4*a*c

	if d>=0 {
		fmt.Println(" Ax^2+bx+c=0 Tenglama  haqiqiy ildzga ega")
	} else {
		fmt.Println("Ax^2+bx+c=0 tenglama diskiriminant nolga teng yoki katta bo'lganda haqiqiy ildizga ega")

	}
	

}