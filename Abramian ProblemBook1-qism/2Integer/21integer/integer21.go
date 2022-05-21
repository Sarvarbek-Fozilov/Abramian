package main

import "fmt"

func main(){
	var n,minut int
	fmt.Println("Kun  boshidan boshlab necha  sekun o'tganini kiritng")
	fmt.Scanf("%d",&n)
	minut=n/60
	fmt.Println("Kun boshidan beri shuncha:",minut,"minut",n-minut*60,"sekund")


}
