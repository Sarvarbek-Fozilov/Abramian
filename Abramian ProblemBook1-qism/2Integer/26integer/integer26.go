package main

import "fmt"

func main(){
	var k int
	fmt.Println("istalgan kundi  sonini kiritng")
	fmt.Scanf("%d",&k)
	x:=k%7+1
	if x==0{
		fmt.Println("Yakshanba")
	}
	if x==1{
		fmt.Println("Dushanba")
	}
	if x==2{
		fmt.Println("Seshanba")
	}
	if x==3{
		fmt.Println("Chorshanba")
	}
	if x==4{
		fmt.Println("Payshanba")
	}
	if x==5{
		fmt.Println("Juma")
	}
	if x==6{
		fmt.Println("Shanba")
	} 

	

}