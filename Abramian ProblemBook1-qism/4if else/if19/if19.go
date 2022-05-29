package main

import "fmt"

func main()  {
	var a,b,c,d int
	fmt.Println("Birinch sondi kiriting")
	fmt.Scanf("%d",&a)
	fmt.Println("Ikkinch sondi kiriting")
	fmt.Scanf("%d",&b)
	fmt.Println("Uchinch sondi kiriting")
	fmt.Scanf("%d",&c)
	fmt.Println("To'rtinch sondi kiriting")
	fmt.Scanf("%d",&d)

	if  a==b && b==c && c==d  {
		fmt.Println("Xamma kiritlgan sonlar o'zaro teng")
	} else if (b==c && c==d){
		fmt.Println("faqat birinchi kiritilgan son qolganlri bilan o'zaro teng emas",a)

	} else if a==c && c==d{
		fmt.Println("faqat ikkinchi kiritilgan son qolganlri bilan o'zaro teng emas",b)

	} else if a==b && b==d {
		fmt.Println("faqat uchinchi kiritilgan son qolganlri bilan o'zaro teng emas",c)
	} else  if  (a==b && b==c ){
		fmt.Println("faqat to'rtinchi kiritilgan son qolganlri bilan o'zaro teng emas",d)
	} else {
		fmt.Println("Uchta o'zaro teng son yo'q")
	}
}