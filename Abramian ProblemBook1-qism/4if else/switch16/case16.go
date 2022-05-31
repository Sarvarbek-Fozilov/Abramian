package main

import "fmt"
func main(){
	var y int
	fmt.Println("Yoshingizni  kirting")
	fmt.Scanf("%d",&y)

	switch y/10{
	case 2:
		fmt.Println("Yigirma")
	case 3:
		fmt.Println("O'ttiz")
	case 4:
		fmt.Println("Qirq")
	case 5:
		fmt.Println("Ellik")
	case 6:
		fmt.Println("Oltmish")    
	}
	switch y%10{
	case 1:
		fmt.Println("bir")
	case 2:
		fmt.Println("ikki")
	case 3:
		fmt.Println("Uch")
	case 4:
		fmt.Println("To'rt")
	case 5:
		fmt.Println("besh")
	case 6:
		fmt.Println("Olti")
	case 7:
		fmt.Println("yetti")
	case 8:
		fmt.Println("Sakkiz")
	case 9:
		fmt.Println("To'qqiz")

	}
	switch (y%10) {
	case 0:
		fmt.Println("yosh.\n")
	case 1:
		fmt.Println("yosh.\n")
	case 2:
		fmt.Println("yosh.\n")
		
	case 3:
		fmt.Println("yosh.\n")
	case 4:
		fmt.Println("yosh.\n")
	case 5:
		fmt.Println("yosh.\n")

	case 6:
		fmt.Println("yosh.\n")
	case 7:
		fmt.Println("yosh.\n")
	case 8:
		fmt.Println("yosh.\n")
	case 9:
		fmt.Println("yosh.\n")
	}	
}