//Umumiy markazga ega bo'lgan ikta aylana radiusi b-gan:R1,R2
//Ularning yuzalari S1 va S2 ularning ayirmasi S3 aniqlansin
//S1=p*r1,S2=pi*R2,S3=pi*(R1-R2)

package main

import "fmt"

func main(){
	var R1,R2,S1,S2,S3,pi float64
	pi=3.14 

	fmt.Println("Aylnananing katta radiusini kiriting")
	fmt.Scanf("\n%d",R1)

	fmt.Println("Aylnananing kichik radiusini kiriting")
	fmt.Scanf("\n%d",R2)

	S1=pi*R1
	S2=pi*R2
	S3=pi*(R1-R2)

	fmt.Println("S1",S1,"ga teng")
	fmt.Println("S1",S2,"ga teng")
	fmt.Println("S1",S3,"ga teng")



}