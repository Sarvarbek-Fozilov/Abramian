package main

import "fmt"
func main(){
	var m int
	fmt.Println("Qaysi masala")
	fmt.Scanf("%d",&m)

	switch m/10{
		case 1:{
			fmt.Println("O'n")
		case 2:
			fmt.Println("Yigirma")
		case 3:
			fmt.Println("O'ttiz")
		case 4:
			fmt.Println("Qirq")
		}
		switch y%10{
		case 0:
			fmt.Println("masala")
		case 1:
			fmt.Println("Bitta masala")
		case 2:
			fmt.Println("ikkita masala")
		case 3:
			fmt.Println("uchta masala")
		case 4:
			fmt.Println("to'rta ,masala")
		case 5:
			fmt.Println("beshta masala")
		case 6:
			fmt.Println("oltita masala")
		case 7:
			fmt.Println("yetita masala")
		case 8:
			fmt.Println("sakkizta masala")
		case 9:
			fmt.Println("to'qizta masala")
		}
	}


}