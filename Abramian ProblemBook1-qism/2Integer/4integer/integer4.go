package main

import "fmt"
func main(){
	var a,b float64
	fmt.Println("A kesmaning uzunligini  musbat sonda kiriting")
	fmt.Scanf("%f",&a)
	fmt.Println("B kesmaning uzunligini  musbat sonda kiriting")
	fmt.Scanf("%f",&b)

	fmt.Println("A kesmada B kesmani ",a/b,"marta joylashtrish mumkin.")
}