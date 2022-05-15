//Aylananing uzunligi l berilgan.
//Uning rasisu r va yzuasi S aniqlansin
// L=2*pi*R,S=pi*r^2

package main

import "fmt"

func main(){
	 var L,S,R float64

	 fmt.Println("Aylananing uzunligini kiriting")
	 fmt.Scanf("%f",&L)
	 pi:=3.14

	 R= L / (2*pi)
	 S= pi*(R*R)

	 fmt.Println("Aylananing Radiusi",R,"ga teng")
	 fmt.Println("Aylananing Yuzasi",S,"ga teng")

}