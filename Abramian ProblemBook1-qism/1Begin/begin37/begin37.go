package main

import (
	"fmt"
	"math"
)
func main(){
	var v1,v2,s,t,s1 float64

	fmt.Println("Birinchi avtomobil tezligini kiriting")
	fmt.Scanf("%f",&v1)
	fmt.Println("ikkinchi avtomobil tezligini kiriting")
	fmt.Scanf("%f",&v2)
	fmt.Println("Ular orasdagi masofani kiriting:")
	fmt.Scanf("%f",&s)

	fmt.Println("Vaqtni kiriting")
	fmt.Scanf("%f",&t)
	s1 = math.Abs(s-(v1*t+v2*t))

	fmt.Println("ikki avtomobil orasdagi masofa",s1)


}