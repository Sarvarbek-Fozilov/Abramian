package main

import "fmt"

func main(){
	var x,y int
	fmt.Println("X oqiga qiymat kirit")
	fmt.Scanf("%d",&x)
	fmt.Println("Y oqiga qiymat kirit")
	fmt.Scanf("%d",&y)

	if x>0 && y>0{
		fmt.Println( "1-chorak")
	} else if x<0 && y>0 {
		fmt.Println("2-chorak")
	}else if x<0 && y<0 {
		fmt.Println("3-chorak")
	} else {
		fmt.Println("4-chorak")
	}
}