package main

import "fmt"

func main(){
	var x,y int
	fmt.Println("X oqiga qiymat kirit")
	fmt.Scanf("%d",&x)
	fmt.Println("Y oqiga qiymat kirit")
	fmt.Scanf("%d",&y)

	if x==0 && y==0 {
		fmt.Println("0")
	} else if  x>=0 || x<=0{
		fmt.Println("1")
	} else if y>=0 || y<=0{
		fmt.Println("2")
	} else {
		fmt.Println("3")
	}


}