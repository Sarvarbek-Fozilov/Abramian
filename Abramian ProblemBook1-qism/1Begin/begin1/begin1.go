//Kvadratning tomoni berilgan a,  uning uning perimetri  aniqlansin P= 4*a

package main

import (
	"fmt"

)

func main(){
	var tomon float64

	fmt.Println("Kvadratning tomonini uzunligini kiriting:")
	fmt.Scanln( &tomon)



	fmt.Println("Kvadratning Perimetri P= 4*a  formulsgs ko'ra P= ", 4*tomon)
}

