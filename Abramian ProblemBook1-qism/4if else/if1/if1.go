package main

import "fmt"


func main(){
	son := 0
	fmt.Println("Istalgan sonizni kiriting")
	fmt.Scanf("%d",&son)
	if son > 0 {
		fmt.Println( son+1)
	}else {
		fmt.Println(son)
	}
}