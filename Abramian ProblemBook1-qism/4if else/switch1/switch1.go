package main

import "fmt"

func main(){
	var day int
	fmt.Println("Kundi kiriting 1dan 7gacha son bilan")
	fmt.Scanf("%d",&day)

	switch day{
	case 1:
		fmt.Println("Dushanba")

	case 2: 
		fmt.Println("Seshanba")

	case 3:
		fmt.Println("Chorshanba")

	case 4:
		fmt.Println("Payshanba")

	case 5:
		fmt.Println("Juma")

	case 6:
		fmt.Println("Shanba")

	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("7gacha bo'lgan son kirting")	
	}
	
}