//Paralapiped yon tomonlari a,b,cberilgan,
//uning hajmini V=a*b*c va tola sirti S=2*(a*b+b*c+a*c)aniqlansin.

package main

import "fmt"


func main(){
	var tomon1,tomon2,tomon3,hajm,sirt float64

	fmt.Println("Paralapipedni birinchi  yon  tomonini  kiriting:")
	fmt.Scanf("%f", &tomon1)

	fmt.Println("Paralapipedni ikkinchi  yon  tomonini  kiriting:")
	fmt.Scanf("%f", &tomon2)

	fmt.Println("Paralapipedni uchinchi yon  tomonini  kiriting:")
	fmt.Scanf("%f", &tomon3)

	hajm= tomon1*tomon2*tomon3
	sirt= 2*(tomon1*tomon2+tomon2*tomon3+tomon1*tomon3)

	fmt.Println("Paralapipedning hajmi:",hajm)
	fmt.Println("Paralapipedning sirti:",sirt)
}