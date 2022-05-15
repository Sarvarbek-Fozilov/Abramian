package main

import (
	"fmt"
	"math"
)
func main(){
	var a,b,c,ac,bc float64

	fmt.Println("A nuqtaning qiymatini kiriting")
	fmt.Scanf("%f",&a)

	fmt.Println("B nuqtaning qiymatini kiriting")
	fmt.Scanf("%f",&b)

	fmt.Println("C nuqtaning qiymatini kiriting")
	fmt.Scanf("%f",&c)

	ac= math.Abs(c - a)
	bc= math.Abs(c-b)
	fmt.Println("AC kesma uzunligi:",ac,"BC kesma uzunligi:",bc,"Kesmalar uzunliklari yig'indisi:",ac-bc)
	

}