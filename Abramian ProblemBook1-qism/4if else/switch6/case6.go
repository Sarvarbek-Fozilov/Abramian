package main

import "fmt"
func main(){
	var metr,uzunlik float64

	fmt.Println("Kesma uzunligini metrda kiriting")
	fmt.Scanf("%f",&metr)
	fmt.Println("O'zgartirmoqchi bo'lgan uzunlik birligini tanlang\n 1.desimetr 2.kilometr 3.metr 4.Millimetr 5.santimetr")
	fmt.Scanf("%f",&uzunlik)

	switch uzunlik{
	case 1:
		fmt.Println("desimetrdagi qiymati:", metr/10 )
	case 2:
		fmt.Println("Kilometrdagi qiymati:",metr*1000 )
	case 3:
		fmt.Println("Metrdagi qiymati:",metr )
	case 4:
		fmt.Println("Millimetrdagi qiymati:",metr/1000 )
	case 5:
		fmt.Println("Santimetrdagi qiymati:",metr/100 )	
	default:
		fmt.Println("To'gri raqamni tanlang")				
	}

}