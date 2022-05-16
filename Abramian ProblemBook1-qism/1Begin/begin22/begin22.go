package main

import "fmt"

func main(){
	var a,b,c float64

	fmt.Println("A sonining qiymatini kiriting")
	fmt.Scanf("%f",&a)
	fmt.Println("B sonining qiymatini kiriting")
	fmt.Scanf("%f",&b)
	c=a
	a=b
	b=c
	fmt.Println("A ning yangi qiymati : ",a)
	fmt.Println("B ning yangi qiymati:",b)
}
