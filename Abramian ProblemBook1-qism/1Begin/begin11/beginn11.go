package main

import (
	"fmt"
	"math"
)
func main(){
	var a,b,c,d float64
	fmt.Println("Birinchi sondi kiriting")
	fmt.Scanf("%f",&a)
	fmt.Println("Ikkinchi sondi kiriting")
	fmt.Scanf("%f",&b)

	c=math.Sqrt(a)
	d=math.Sqrt(b)

	fmt.Println("Ularning yigindisi:",a+b)
	fmt.Println("Ularning ko'paytmasi:",a*b)
	fmt.Println("birinchi sonning moduli:",c)
	fmt.Println("Ikkinchi sonning moduli:",d)
}