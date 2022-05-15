//Ikkta manfiy son berilgan
// Ularning o'rta geometrigini toping.

package main

import (
    "fmt"
    "math"
)


func main( ){
	 var a,b, c, float64

	 fmt.Println("Birinchi sondi kiriting")
	 fmt.Scanf("%f",&a)

	 fmt.Println("Ikkinchi sondi kiriting")
	 fmt.Scanf("%f",&b)

	 c=a*b

	 res:= math.Sqrt(c)

	 fmt.Println("O'rta geometrigi", res,"ga teng")

}