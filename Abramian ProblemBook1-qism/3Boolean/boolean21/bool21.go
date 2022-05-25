package main

import "fmt"

func main(){
	var son,yuzlik,birlik,onlik int

	fmt.Println("Uchxonalik sondi kiriting")
	fmt.Scanf("%d",&son)

	yuzlik= son/100
	onlik= (son- yuzlik * 100)/10
	birlik= (son- yuzlik * 100)- onlik*10

	if yuzlik<onlik && onlik < birlik {
		fmt.Println("Sonning raqamlari ketma-ket o'suvchi")
	} else{
		fmt.Println("Boshqa son kiriting")
	}

}