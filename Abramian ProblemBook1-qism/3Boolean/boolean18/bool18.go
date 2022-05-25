package main

import "fmt"

func main(){
	var a,b,c int
	fmt.Println("Birinchi sondi kiritning")
	fmt.Scanf("%d",&a)
	fmt.Println("Ikkinchi sondi kiritning")
	fmt.Scanf("%d",&b)
	fmt.Println("Uchinchi sondi kiritning")
	fmt.Scanf("%d",&c)

	if a==b || b==c ||a==c{
		fmt.Println("Berilgn sondan xech bo'lmaganda 2tasi bir biriga teng")
	} else {
		fmt.Println(" teng son yo'q")
	}
}