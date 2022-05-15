//To'g'ri to'rtburchakning  tomonlari a va b berilgan,
//  uning yuzasi va perimetri aniqlansin S= a*b,P = 2*(a+b)

package main

import "fmt"

func main(){
	var tomon1,tomon2,yuza, perimetr float64

	fmt.Println("To'gri to'rtburchakning birinchi tomonini uzunligini kiriting:")
	fmt.Scanf("%f", &tomon1)

	fmt.Println("To'gri to'rtburchakning ikkinchi tomonini uzunligini kiriting:")
	fmt.Scanf("%f", &tomon2)

	perimetr = 2*(tomon1+tomon2)
	yuza = tomon1*tomon2

	fmt.Println("To'g'ri to'rtburchakning Perimetri P= 2*(a+b) formulaga ko'ra P=", perimetr)
	fmt.Println("To'g'ri to'rtburchskning Yuzi S=a*b formulaga ko'ra S=",yuza)
}

