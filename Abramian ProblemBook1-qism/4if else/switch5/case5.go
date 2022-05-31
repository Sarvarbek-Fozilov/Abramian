package main

import "fmt"

func main(){
	var a,b,amal float64
	fmt.Println("A sonini kiriting")
	fmt.Scanf("%f",&a)

	fmt.Println("B sonini kiriting")
	fmt.Scanf("%f",&b)

	fmt.Println("Amalni tanlang!:\n 1.qo'shish  2.ayrish 3.bo'lish 4.Ko'paytirish")
	fmt.Scanf("%f",&amal)

	switch amal {
	case 1:
		fmt.Println("Yig'indi:",a+b)
	case 2:
		fmt.Println("Ayirma:",a-b)	
	case 3:
		fmt.Println("Bo'linma:", a/b)
	case 4:
		fmt.Println("Ko'paytma:",a*b)
	default:
		fmt.Println("To'gri amal tanla")			
	}


}