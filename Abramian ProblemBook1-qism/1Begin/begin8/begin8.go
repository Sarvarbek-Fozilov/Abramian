// ikkta son a va b berigan
// ularning o'rta arifmetigini toping (a+b)/2

package main

import "fmt"

func main(){
	var a,b,arifmetik float32
	 fmt.Println("Birinchi sondi kiriting:")
	 fmt.Scanf("%d" , &a)

	 fmt.Println("Ikkinchi sondi kiriting:")
	 fmt.Scanf("%d" , &b)

	 arifmetik = (a+b)/2

	 fmt.Println("O'rta arifmetigi",arifmetik,"ga teng")
}