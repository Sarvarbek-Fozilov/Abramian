package main

import "fmt"

func main(){
	var x1,y1,x2,y2 int

	fmt.Println("Shaxmat doskasining X1 kordintasi qiymatni kirit")
	fmt.Scanf("%d",&x1)
	fmt.Println("Shaxmat doskasining Y1 kordintasi qiymatni kirit")
	fmt.Scanf("%d",&y1)

	fmt.Println("Shaxmat doskasining X2 kordintasi qiymatni kirit")
	fmt.Scanf("%d",&x2)
	fmt.Println("Shaxmat doskasining Y2  kordintasi qiymatni kirit")
	fmt.Scanf("%d",&y2)
	

	if x1==x2|| y1==y2 {
		fmt.Printf("Tog'ri\n")
	} else {
		fmt.Printf("notog'ri\n")
	}
}