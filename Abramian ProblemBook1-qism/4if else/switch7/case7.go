package main

import "fmt"
func main(){
	var massa,ogirlik float64

    fmt.Println("Bilmoqchi bo'lgan massa qiymatini kiriting")
	fmt.Scanf("%d",&massa)
	fmt.Println("Bilmoqchi bo'lgan og'irlik birligini tanlang\n 1.kilogramm 2.milligram 3.gramm 4.tonna 5.sentner")
	fmt.Scanf("%f",&ogirlik)

	switch ogirlik{
	case 1:
		fmt.Println("Kilogramdagi qiymati:", massa )
	case 2:
		fmt.Println("Milligramdagi qiymati:",massa*1000000 )
	case 3:
		fmt.Println("gramdagi qiymati:",massa/1000 )
	case 4:
		fmt.Println("Tonnadagi qiymati:",massa/1000 )
	case 5:
		fmt.Println("Sentnerdagi qiymati:",massa/100 )	
	default:
		fmt.Println("To'gri raqamni tanlang")				
	}

}