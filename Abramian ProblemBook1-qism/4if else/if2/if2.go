package main

import "fmt"


func main(){
	son := 0
	fmt.Println("Istalgan sonizni kiriting")
	fmt.Scanf("%d",&son)
	if son > 0 {
		fmt.Printf("Musbat%d", son+1)
	}else {
		fmt.Println(son-2)
	}
}