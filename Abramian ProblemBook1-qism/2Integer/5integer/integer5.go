package main

import "fmt"
func main(){
	var a,b int
	fmt.Println("A kesmaning uzunligini  musbat sonda kiriting")
	fmt.Scanf("%d",&a)
	fmt.Println("B kesmaning uzunligini  musbat sonda kiriting")
	fmt.Scanf("%d",&b)

	fmt.Println("A kesmada B kesmani ",a/b,"marta joylashtrish mumkin.")
	fmt.Println("A kesmada b kesmaning",a-(a/b)*b,"joylashmagan qismi")
}