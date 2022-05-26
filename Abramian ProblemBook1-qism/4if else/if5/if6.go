package main

import "fmt"

func main() {
	var son1,son2,son3,mus,man int

	fmt.Println("Birinchi sondi kiriting")
	fmt.Scanf("%d", &son1)
	fmt.Println("Ikkinchi sondi kiriting")
	fmt.Scanf("%d", &son2)
	fmt.Println("Uchinchi sondi kiriting")
	fmt.Scanf("%d", &son3)

	if (son1>0) {
		mus++;

	}else {
		man=1;
	} 
   
	if (son2>0) {
		mus++;

	}else {
		man++;
	} 
   
	if (son3>0) {
		mus++;

	}else {
		man++;
	} 
   
	fmt.Printf("Sonlar orasida  %d  musbat  %d manfiy sonlar bor\n" ,mus, man);
}