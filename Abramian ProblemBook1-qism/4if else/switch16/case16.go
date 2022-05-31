package main

import "fmt"
func main(){
	var y int
	fmt.Println("Yoshingizni  kirting")
	fmt.Scanf("%d",&y)

	switch y/10{
	case 1:
		fmt.Printf("O'n ")	
	case 2:
		fmt.Printf("Yigirma ")
	case 3:
		fmt.Printf("O'ttiz ")
	case 4:
		fmt.Printf("Qirq ")
	case 5:
		fmt.Printf("Ellik ")
	case 6:
		fmt.Printf("Oltmish ")    
	}
	switch y%10{
	case 1:
		fmt.Printf("bir ")
	case 2:
		fmt.Printf("ikki ")
	case 3:
		fmt.Printf("Uch ")
	case 4:
		fmt.Printf("To'rt ")
	case 5:
		fmt.Printf("besh ")
	case 6:
		fmt.Printf("Olti ")
	case 7:
		fmt.Printf("yetti ")
	case 8:
		fmt.Printf("Sakkiz ")
	case 9:
		fmt.Printf("To'qqiz ")

	}
	switch (y%10) {
	case 0:
		fmt.Printf("yosh.\n")
	case 1:
		fmt.Println("yosh.")
	case 2:
		fmt.Println("yosh.")
		
	case 3:
		fmt.Println("yosh.")
	case 4:
		fmt.Println("yosh.")
	case 5:
		fmt.Println("yosh.")

	case 6:
		fmt.Println("yosh.")
	case 7:
		fmt.Println("yosh.")
	case 8:
		fmt.Println("yosh.")
	case 9:
		fmt.Println("yosh.")
	}	
}