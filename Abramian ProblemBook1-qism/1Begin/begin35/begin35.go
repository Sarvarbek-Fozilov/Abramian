package main

import "fmt"
func main(){
	var v,u,t1,t2,s float64

	fmt.Println("Qayiqning tezligini kiriting")
	fmt.Scanf("%f",&v)
	fmt.Println("Daryo oqimining tezligi:")
	fmt.Scanf("%f",&u)

	fmt.Println("Qayiqning daryo oqimi bo'yicha harakatlanish vaqti:")
	fmt.Scanf("%f",&t1)

	fmt.Println("Qayiqning daryo oqimiga qarshi harakatlanish vaqti:")
	fmt.Scanf("%f",&t2)
	s=v*t1+(v-u)*t2
	fmt.Println("Qayiqni yurgan yo'li:",s)


}