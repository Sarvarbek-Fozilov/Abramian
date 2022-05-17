package main

import (
	"fmt"
	"math"
)

func main(){
	var x,y float64

	fmt.Println("X ning qiymatini kiritig")
	fmt.Scanf("%f",&x)

	y=4*(math.Pow((x-3),6)-7*(math.Pow((x-3),3))+2)
	y=4*(math.Pow(x-3,6))-7*(math.Pow(x-3,2))+2

	fmt.Println("Y ning qiymati: ",y)
}