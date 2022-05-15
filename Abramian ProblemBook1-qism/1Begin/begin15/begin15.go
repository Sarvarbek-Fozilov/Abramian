//Aylananing yuzasi S berilganUning diametri d va radiusi r aniqlansin
//S = πr^2, d=2r

package main

import (
	"fmt"
	"math"
)

func main(){
	var S,d,r float64

	fmt.Println("Aylananing yuzini kiriting")
	fmt.Scanf("%f",&S)

	pi:=3.14

	r= math.Sqrt(S/pi)

	d=2*r
	fmt.Println("Aylana diametri:",d)
	fmt.Println("Aylanan radiusi:",r)
}