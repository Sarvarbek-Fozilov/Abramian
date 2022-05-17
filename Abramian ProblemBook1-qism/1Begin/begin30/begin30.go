package main

import "fmt"

func main(){
	var a float64
	fmt.Println("radian qiymatini kirting")
	fmt.Scanf("%f",&a)

	degree:=(180*a)/3.14
	fmt.Println("RAdianning gradus qiymati:",degree)
}