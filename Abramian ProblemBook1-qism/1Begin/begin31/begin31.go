package main

import "fmt"

func main(){
	var  t int
	fmt.Println("Gradusda temperaturani kiriting")
	fmt.Scanf("%d",&t)

	c:=((t-32)*5)/9
	fmt.Println("Gradusni frangeytga o'tkazish:",c)
}