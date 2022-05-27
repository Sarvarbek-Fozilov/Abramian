package main

import "fmt"

func main(){
	var x,y int

	fmt.Println("Birinchi butun sondi kiriting")
	fmt.Scanf("%d",&x)

	fmt.Println("Ikkinnchi butun sondi kiriting")
	fmt.Scanf("%d",&y)

	if  x<y {
		fmt.Println(y,x)
	}

}