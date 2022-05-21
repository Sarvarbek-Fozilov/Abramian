package main

import "fmt"
func main(){
	var kb float64

	fmt.Println("Necha baytni kb da bilmoqchisiz")
	fmt.Scanf("%d",&kb)

	fmt.Println("Ushbu fayl",kb/1024,"kbga teng")
}