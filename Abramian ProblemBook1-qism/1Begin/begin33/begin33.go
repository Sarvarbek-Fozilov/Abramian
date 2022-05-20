package main

import "fmt"
func main(){
	var x,a,y int
	fmt.Println("Necha keg kanfet olmoqchisz")
	fmt.scanf("%d",&x)
	fmt.Println("Necha pullik kanfetdan olmoqchisz")
	fmt.Scanf("%d",&a)
	fmt.Println("necha kg kanfet yana olasz")
	fmt.Scanf("%d",&y)

	price:=a/x
	fmt.Println("narx: ",price)
	fmt.Println("Umumiy qiymat: ",y*price)
}