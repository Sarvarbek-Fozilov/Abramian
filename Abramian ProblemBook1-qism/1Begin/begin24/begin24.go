package main

import "fmt"
func main(){
	var a,b,c,d float64

	fmt.Println("A sonining qiymatini kiriting:",a)
	fmt.Scanf("%f",&a)
	fmt.Println("B sonining qiymatini kiriting:",b)
	fmt.Scanf("%f",&b)
	fmt.Println("C sonining qiymatini kiriting:",c)
	fmt.Scanf("%f",&c)

	d=a
	a=c
	c=b
	b=d
	fmt.Println("A ning yangi qiymati:",a)
	fmt.Println("B ning yangi qiymati:",b)
	fmt.Println("C ning yangi qiymati:",c)


}