package main

import "fmt"

func main(){
	var a,b,c float64
	fmt.Println("Sonlar o'qidagi Birinchi nuqtani kiriting")
	fmt.Scanf("%f",&a)
	fmt.Println("Sonlar o'qidagi ikkinchi nuqtani kiriting")
	fmt.Scanf("%f",&b)
	fmt.Println("Sonlar o'qidagi uchinchi nuqtani kiriting")
	fmt.Scanf("%f",&c)

	if a>b && b>c{
		fmt.Println("Masofa",a-b)
	} else if a>c && c>b{
		 fmt.Println("Masofa",a-c)
	} else if a<b && b<c{
		fmt.Println("Masofa",b-a)

	} else if  a<c && c<b{
		fmt.Println("Masofa",c-a)

	} else if a==b || b==c {
		fmt.Println("Masofa",c-a)
	} else {
		fmt.Println("boshqatdan kirgiz")
	}
}