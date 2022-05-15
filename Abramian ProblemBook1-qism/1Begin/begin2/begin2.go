//Kvadratning tomoni berilgan a,  uning uning  yuzini aniqlang S= a^2

package main

import "fmt"

func main(){
	var tomon, yuza float64

	fmt.Println("Kvadratning tomonini uzunligini kiriting:")
	fmt.Scanf("%f", &tomon)

	yuza = tomon * tomon

	fmt.Println("Kvadratning Yuzi S=a*a formulaga ko'ra S=", yuza)
}

