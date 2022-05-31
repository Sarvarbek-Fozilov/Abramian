package main

import "fmt"

func main(){
	var k int
	fmt.Println("Baxoyingizni kiriting")
	fmt.Scanf("%d",&k)
	switch k{
	case 1:
		fmt.Println("1 ball bu yomon")
	case 2:
		fmt.Println("2 ball bu qoniqarsiz!")
	case 3:
		fmt.Println("3 ball bu qoniqarli")
	case 4:
		fmt.Println("4 ball bu yaxshi!")
	case 5:
		fmt.Println("5 ball bu A'lo")
	default:
		fmt.Println("Xato")					
	}
}