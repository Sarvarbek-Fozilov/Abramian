package main

import "fmt"

func main(){
	var x,y ,x1,y1,x2,y2 float64
	fmt.Println("Kordinataning X o'qiga qiymat bering")
	fmt.Scanf("%f",&x)

	fmt.Println("Kordinataning Y o'qiga qiymat bering")
	fmt.Scanf("%f", &y)
	fmt.Println("Kordinataning X1 o'qiga qiymat bering")
	fmt.Scanf("%f",&x1)

	fmt.Println("Kordinataning Y1 o'qiga qiymat bering")
	fmt.Scanf("%f", &y1)
	fmt.Println("Kordinataning X2 o'qiga qiymat bering")
	fmt.Scanf("%f",&x2)

	fmt.Println("Kordinataning Y2 o'qiga qiymat bering")
	fmt.Scanf("%f", &y2)

	if x1<x && x<x2 && y2<y && y<y1 {
		fmt.Println(" tomonlar Kordinata o'qlariga parallel bo'lgan to'rtburchak ichida yotadi")
	} else{
		fmt.Println("Qayta urunb ko'ring")
	}
}