package main

import "fmt"
func main(){
	var son,yuzlik,onlik,birlik int
	fmt.Println("uch xonalik son kiriting")
	fmt.Scanf("%d",&son)
	yuzlik=son/100
	onlik=(son-yuzlik*100)/10
	birlik=son-yuzlik*100 - onlik*10
	fmt.Println("Ushbu uch xonalik sondi sonlar yig'indisi:",yuzlik+onlik+birlik)



}