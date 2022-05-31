package main

import "fmt"

func main(){
	var num int

	fmt.Println("Son kiritng")
	fmt.Scanf("%d",&num)

		 if num>999 {
			fmt.Println("Uchxonalikdan ham katta son kirtdingiz")
		} else if num>99 && num<=999 {
			fmt.Println("Uch xonalik son")
		} else if num>9 && num<=99 {
			fmt.Println("Ikki xonalik son")
		} else {
			fmt.Println("Qayta kiting")
		}
	

}