package main

import "fmt"

func main(){
	var  fasl int
	fmt.Println("Oy raqamini kiriting")
	fmt.Scanf("%d",&fasl)

	switch fasl{
	case 1:
		fmt.Println("Qish fasli")
	case 2:
		fmt.Println("Qish fasli")
	case 3:
		fmt.Println("Baxor fasli")	
	case 4:
		fmt.Println("Baxor fasli")	
	case 5:
		fmt.Println("Baxor fasli")	
	
	case 6:
		fmt.Println("Yoz fasli")	
	case 7:
		fmt.Println("Yoz fasli")	
	case 8:
		fmt.Println("Yoz fasli")

	case 9:
		fmt.Println("Kuz fasli")	
	case 10:
		fmt.Println("Kuz fasli")	
	case 11:
		fmt.Println("Kuz fasli")
	case 12:
		fmt.Println("Qish fasli")
	default:
		fmt.Println("To'g'ri kirit")
	
	}
}