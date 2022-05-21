package main

import "fmt"
func main(){
	var year int
	fmt.Println("Yilni kiriting uni nechinchi asrga tegishlik ekanligini  bilib oling")
	fmt.Scanf("%d",&year)

	fmt.Println("Kiritilgan yil:",year/100+1,"-asrga teng")
}