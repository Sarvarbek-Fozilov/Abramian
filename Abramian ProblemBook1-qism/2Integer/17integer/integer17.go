package main

import "fmt"
func main(){
	var son,yuzlik,minglik int
	fmt.Println("tort xonalik son kiriting")
	fmt.Scanf("%d",&son)

	minglik=son/1000
	yuzlik=(son-minglik*1000)/100
	fmt.Println("Kiritilgan sonning yuzlik soni :",yuzlik)


}