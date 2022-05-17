package main

import (
	"fmt"
	"math"
)

func main(){
	var a,a2,a3,a5,a10,a15 float64

	fmt.Println("A sonining qiymatini kiriting")
	fmt.Scanf("%f",&a)

	a2=math.Pow(a,2)
	a3=math.Pow(a,3)
	a5=math.Pow(a,5)
	a10=math.Pow(a,10)
	a15=math.Pow(a,15)

	fmt.Println("A ning 2-darajasi",a2)
	fmt.Println("A ning 3-darajasi",a3)
	fmt.Println("A ning 5-darajasi",a5)
	fmt.Println("A ning 10-darajasi",a10)
	fmt.Println("A ning 15-darajasi",a15)
}