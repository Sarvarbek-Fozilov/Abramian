package main

import "fmt"

func main(){
	var son1,son2 int
	fmt.Println("Birinchi sondi kiriting")
	fmt.Scanf("%d", &son1)
	fmt.Println("Ikkinchi sondi kiriting")
	fmt.Scanf("%d", &son2)

	if son1 > son2{
		fmt.Println("Birnchi son katta")
	} else if son1 <son2 {
		fmt.Println("ikkinchi sondi kirting")
		
	} else {
		fmt.Println("Boshqa son kirting")
	}
	
}