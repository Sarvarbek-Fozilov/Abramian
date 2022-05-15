// Aylananing diametri d berilgan Uning uzunligi aniqlansin L= pi*d, pi=3.14

package main

import "fmt"


func main(){
	var diametr,p float32
	p = 3.14

	fmt.Println("Aylananing diametrini kiriting:")
	fmt.Scanf("%v",diametr)

	uzunlik := diametr * p
	

	fmt.Println("Aylananing uzunligi L=pi*d formulaga ko'ra L=", uzunlik)
} intel inside 
4/32gb