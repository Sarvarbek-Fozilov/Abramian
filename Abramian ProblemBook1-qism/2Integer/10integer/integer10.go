package main

import "fmt"
func main(){
	var son,yuzlik,onlik  int
	fmt.Println("uch xonalik son kiriting")
	fmt.Scanf("%d",&son)
	yuzlik=son/100
	onlik=(son-yuzlik*100)/10
	fmt.Println("yuzliklar xonasdagi son:",yuzlik)
	fmt.Println("oniklar xonasdagi son:",onlik)


}