package main

import "fmt"

func main(){
	var k int
	fmt.Println("Kundi kiriting: ")
	fmt.Scanf("%d",&k)

	switch (k+3)%7{
	case 0:
		fmt.Println("Yakshanba")
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
	default:
		fmt.Println("Adashdinooov")			
	}

}