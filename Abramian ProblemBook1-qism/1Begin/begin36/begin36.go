package main

import "fmt"
func main(){
	var v1,v2,s,t float64

	fmt.Println("Birinchi avtomobil tezligini kiriting")
	fmt.Scanf("%f",&v1)
	fmt.Println("ikkinchi avtomobil tezligini kiriting")
	fmt.Scanf("%f",&v2)
	fmt.Println("Ular orasdagi masofani kiriting:")
	fmt.Scanf("%f",&s)

	fmt.Println("Vaqtni kiriting")
	fmt.Scanf("%f",&t)

	fmt.Println("ikki avtomobil orasdagi masofa",s+v1*t+v2*t)


}