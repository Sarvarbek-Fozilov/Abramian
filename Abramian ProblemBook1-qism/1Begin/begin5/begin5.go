//Kubning yon tomoni berilgan,
//uning hajmini V=a^3 va tolsa sirti S=6*a^2 aniqlansin.

package main

import "fmt"


func main(){
	var tomon,hajm,sirt float64

	fmt.Println("Kubning yon tomonini  kiriting:")
	fmt.Scanf("%f", &tomon)

	hajm= tomon*tomon*tomon
	sirt= 6*(tomon*tomon)

	fmt.Println("Aylananing hajmi:",hajm)
	fmt.Println("Aylananing sirti:",sirt)
}