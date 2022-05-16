package main

import (
	"fmt"
	"math"
)

func main(){
	var x,y float64

	fmt.Println("X ning qiymatini kiriting")
	fmt.Scanf("%f",&x)

	y=3*(math.Pow(x,6)-6*(math.Pow(x,2))-7)

	fmt.Println("Y ning qiymati:",y)
}