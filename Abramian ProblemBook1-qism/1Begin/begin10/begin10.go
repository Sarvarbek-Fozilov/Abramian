//Nolga teng bo'lmagan ikkta son berilgan
//Ularning yig'indisni,ko'paytmasini, va  har birining kvadrati aniqlansin

package main

import (
	"fmt"

)

func main(){
	var a,b,c,d, yigindi,kopaytma float64

	fmt.Println("Birinchi sondi kiriting")
	fmt.Scanf("%f",&a)

	fmt.Println("Ikkinchi sondi kiriting")
	fmt.Scanf("%f",&b)

	yigindi=(a+b)
	kopaytma=a*b
	c=a*a
	d =b*b

	fmt.Println("Yig'indisi", yigindi ,"ga teng")
	fmt.Println("Ko'paytmasi", kopaytma ,"ga teng")
	fmt.Println("Birinchi sondi kvadrati:",c ,"ga teng")
	fmt.Println("Ikkinchi sondi kvadrati:", d ,"ga teng")

}