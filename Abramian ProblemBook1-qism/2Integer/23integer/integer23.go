package main

import "fmt"

func main(){
	var n,minut,soat,sekund int
	fmt.Println("Kun  boshidan boshlab necha  sekun o'tganini kiritng")
	fmt.Scanf("%d",&n)

	soat=n/3600
	minut=(n-soat*3600)/60
	sekund=n-soat*3600-minut*60
	fmt.Println("Kun boshidan beri shuncha:",soat,"soat",minut,"minut",sekund,"sekund")


}
