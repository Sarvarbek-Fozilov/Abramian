package main

import "fmt"

func main(){
	var a,b,c float64

	fmt.Println("Birinchi sondi kiriting")
	fmt.Scanf("%f",&a)
	fmt.Println("Ikkinchi sondi kiriting")
	fmt.Scanf("%f",&b)
	fmt.Println("Uchinchi sondi kiriting")
	fmt.Scanf("%f",&c)

	if a==b && b==c {
		fmt.Println("Barchasi o'zaro teng")
	} else if a==b{
		fmt.Println("Birinchi va ikkinchi kiritilgan sonlar o'zaro teng  uchinchi son esa bu:",c)
	} else if b==c {
		fmt.Println("ikkinchi va uchinchi kiritilgan sonlar ozaro teng, bu esa birinchi son",a)
	} else if a==c {
		fmt.Println("Birinchi va uchinchi kiritilgan sonlar o'zaro teng, bu esa ikkinchi kiritlgan son",b)

	} else {
		fmt.Println("o'zaro teng sonlar mavjud emas")
	}
}