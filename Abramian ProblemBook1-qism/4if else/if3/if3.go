package main

import "fmt"


func main(){
	son := 0
	fmt.Println("Istalgan sonizni kiriting")
	fmt.Scanf("%d",&son)
	if son > 0 {
		fmt.Println(son+1)
	}else if son<0 {
		fmt.Println(son-2)
	}else if  son==0 {
		fmt.Println(son +10)


	}
}