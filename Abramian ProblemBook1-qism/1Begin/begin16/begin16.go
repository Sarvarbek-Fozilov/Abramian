package main

import (
	"fmt"
	"math"
)

func main() {
    var a,b float64
	
    fmt.Println("Birinchi sondi kiriting:")
	fmt.Scanf("%f\n",&a)

	fmt.Println("Ikkinchi sondi kiriting:")
	fmt.Scanf("%f\n",&b)
	x := math.Abs(a-b)
	fmt.Printf("%.1f\n", x)

}