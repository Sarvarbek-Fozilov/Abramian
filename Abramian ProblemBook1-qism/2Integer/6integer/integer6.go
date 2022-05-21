package main

import "fmt"
func main(){
	var son,onlik,birlik int
	fmt.Println("Ikki xonalik sondi kiriting")
	fmt.Scanf("%d",&son)

	onlik=son/10
	birlik=son-onlik*10
	fmt.Println("Kiritilgan ikkixonalik sonning o'nlik qismi",onlik)
	fmt.Println("Kiritilgan ikkixonalik sonning birik qismi",birlik)
}