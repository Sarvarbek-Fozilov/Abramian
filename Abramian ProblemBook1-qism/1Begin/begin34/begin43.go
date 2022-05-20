package main

import "fmt"

func main(){
	var x,y,a,b int

	fmt.Println("Necha kg Shokolad olmoqchisiz")
	fmt.Scanf("%d",&x)
	fmt.Println("Necha pullik Shokolad dan sotib olmoqchisiz")
	fmt.Scanf("%d",&a)

	fmt.Println("Yana necha kg kanfet olmoqchisiz")
	fmt.Scanf("%d",&y)
	fmt.Println("Bu safar necha pullik kanfetdan sotib olmoqchisiz")
	fmt.Scanf("%d",&b)

	price1:=a/x
	price2:=b/y


	fmt.Println("Shokolad qanfatddan shuncha  barobar qimmat turadi:",price1/price2)


}