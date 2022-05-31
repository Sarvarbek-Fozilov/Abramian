package main

import "fmt"

func main(){
	var n,m int
	fmt.Println("Karta qiymatini kiriting")
	fmt.Scanf("%d",&n)
	fmt.Println("Karta turini kiritng")
	fmt.Scanf("%d",&m)

	switch n{
	case 6:
		fmt.Println("Olti")
	case 7:
		fmt.Println("yetti")
	case 8:
		fmt.Println("Sakkiz")
	case 9:
		fmt.Println("To'qqiz")
	case 10:
		fmt.Println("O'n")
	case 11:
		fmt.Println("valet")
	case 12:
		fmt.Println("Dama")
	case 13:
		fmt.Println("Karol")
	case 14:
		fmt.Println("Tuz")

	}
	switch m{
	case 1:
		fmt.Println("G'ishtim")
	case 2:
		fmt.Println("Yurak")
	case 3:
		fmt.Println("Chilla")
	case 4:
		fmt.Println("Qarg'a")
	}

	

}