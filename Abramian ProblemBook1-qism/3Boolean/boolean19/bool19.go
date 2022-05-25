package main

import "fmt"

func main(){
	var a,b,c int
	fmt.Println("Birinchi sondi kiritning")
	fmt.Scanf("%d",&a)
	fmt.Println("Ikkinchi sondi kiritning")
	fmt.Scanf("%d",&b)
	fmt.Println("Uchinchi sondi kiritning")
	fmt.Scanf("%d",&c)

	if (a==(-1)*b) || (a==(-1)*c) || (b==(-1)*c){
		fmt.Println("Berilgan sonlarning xech bo'lmaganda bir jufti o'zaro qarama-qarshi")

	} else {
		fmt.Println("qayta urnb koring")
	}  
}