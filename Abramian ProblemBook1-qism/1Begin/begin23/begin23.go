package main

import "fmt"
func main(){
	var a,b,c,d float64
	fmt.Println("A sonini qiymatini kiriting")
	fmt.Scanf("%f",&a)
	fmt.Println("B sonini qiymatini kiriting")
	fmt.Scanf("%f",&b)
	fmt.Println("C sonini qiymatini kiriting")
	fmt.Scanf("%f",&c)
	d=a
	a=b
	b=c
	c=d
	fmt.Println("A ning yangi qiymati:",a)
	fmt.Println("B ning yangi qiymati:",b)
	fmt.Println("C ning yangi qiymati:",c)



}