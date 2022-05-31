package main

import "fmt"
func main(){
	var month int
	fmt.Println("Oy raqamini kiriting")
	fmt.Scanf("%d",&month)
	switch (month) {
	case 1:
		fmt.Println("Bu oy 31 kundan iborat")
	case 3:
		fmt.Println("Bu oy 31 kundan iborat")
	case 5:
		fmt.Println("Bu oy 31 kundan iborat")
	case 7:
		fmt.Println("Bu oy 31 kundan iborat")
	case 8:
		fmt.Println("Bu oy 31 kundan iborat")
	case 10:
		fmt.Println("Bu oy 31 kundan iborat")
	case 12:
		fmt.Println("Bu oy 31 kundan iborat")

	case 4:
		fmt.Println("Bu oy 30 kundan iborat")
	case 6:
		fmt.Println("Bu oy 30 kundan iborat")
	case 9:
		fmt.Println("Bu oy 30 kundan iborat")
	case 11:
		fmt.Println("Bu oy 30 kundan iborat")
	case 2:
		fmt.Println("Bu oy 28 kundan iborat")
	default:
		fmt.Println("Oy raqamini to'g'ri kiriting")	
	}
}