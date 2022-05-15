//Doiraning radiusi R berilgan Uning uzunligi va Yuzasi  aniqlansin
// L= 2*pi*R, S=pi*r^2

package main

import "fmt"


func main(){
	var R,L,S float64
	pi := 3.14

	fmt.Println("Doiraning Radiusini kiriting:")
	fmt.Scanf("%v",&R)

	L = 2*pi*R
	S = pi*(R*R)
	

	fmt.Println("Doiraning uzunligi L=", L)
	fmt.Println("Doiraning Yuzasi S=", S)
}