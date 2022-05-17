package main

import (
	"fmt"
	"math"
)

func main(){
	var a,a2,a4,a8 float64

	fmt.Println("A sonining qiymatini kiriting")
	fmt.Scanf("%f",&a)

	a2=math.Pow(a,2)
	a4=math.Pow(a,4)
	a8=math.Pow(a,8)

	fmt.Println("A ning 2-darajasi",a2)
	fmt.Println("A ning 4-darajasi",a4)
	fmt.Println("A ning 8-darajasi",a8)
}