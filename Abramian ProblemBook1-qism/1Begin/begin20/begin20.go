//Tekislikda berilgan ikki nuqta(x1,y1)va (x2,y2) orasdgai masoda aniqlansin
// umumiy ildiz ostida(x2-x1)^2+(y2-y2)^2

package main

import (
	"fmt"
	"math"
)

func main(){
	 var  x1,x2,y1,y2,m float64
	

	 fmt.Println("Birinchi nuqtaning x o'qidagi qiymatini kiriting")
	 fmt.Scanf("%f",&x1)
	 fmt.Println("Birinchi nuqtaning y o'qidagi qiymatini kiriting")
	 fmt.Scanf("%f",&y1)

	 fmt.Println("Ikkinchi nuqtaning x o'qidagi qiymatini kiriting")
	 fmt.Scanf("%f",&x2)

	 fmt.Println("Ikkinchi nuqtaning y o'qidagi qiymatini kiriting")
	 fmt.Scanf("%f",&y2)

	 m =(x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)

	 res:= math.Sqrt(m)
	 fmt.Println("Masofa",res,"ga teng")
}
