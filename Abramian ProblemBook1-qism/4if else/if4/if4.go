package main

import "fmt"

func main() {
	var son1,son2,son3 int

	fmt.Println("Birinchi sondi kiriting")
	fmt.Scanf("%d", &son1)
	fmt.Println("Ikkinchi sondi kiriting")
	fmt.Scanf("%d", &son2)
	fmt.Println("Uchinchi sondi kiriting")
	fmt.Scanf("%d", &son3)

	if (son1>0) {
		son1=1;

	}else {
		son1=0;
	} 
   
	if (son2>0) {
		son2=1;

	}else {
		son2=0;
	} 
   
	if (son3>0) {
		son3=1;

	}else {
		son3=0;
	} 
   
	fmt.Printf("%d\n",son1+son2+son3);
}