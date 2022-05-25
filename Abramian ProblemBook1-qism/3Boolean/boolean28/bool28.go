package main

import "fmt"

func main(){
	var x,y float64
	fmt.Println("Kordinataning x o'qiga qiymat bering")
	fmt.Scanf("%f",&x)

	fmt.Println("Kordinataning y o'qiga qiymat bering")
	fmt.Scanf("%f", &y)

	if (x>0 && y>0) || x<0 && y<0 {
		fmt.Println("Kordinata choragining Birinchi yoki Uchinchisda yotadi.")
	} else{
		fmt.Println("Kordinata choragining Birinchi yoki uchinchisdan tawqari qolganlaridan birida yotadi.")
	}
}