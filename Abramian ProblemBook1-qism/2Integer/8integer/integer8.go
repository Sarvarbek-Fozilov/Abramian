package main

import "fmt"
func main(){
	var son,birlik,onlik int
	fmt.Println("Ikki xonalik sondi kiriting")

	fmt.Scanf("%d",&son)

	onlik=son/10
	birlik=son-onlik*10



	fmt.Println("Kiritilgan sonning  o'rni almashgan xolati:",birlik*10+onlik)
}