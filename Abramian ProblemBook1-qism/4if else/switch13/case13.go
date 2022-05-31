package main

import (
	"fmt"
	"math"
)

func main(){
	var x,a,c,h,s float64
	var nomer int

	fmt.Println("Element nomerini tanlang\n 1.katet 2.gipotenuza 3.gipotenuzaga tushirlgan balandlik 4.Yuzasi")
	fmt.Scanf("%d",&nomer)
	fmt.Println("Uchburchakning bir tomonini uzunligini kiriting")
	fmt.Scanf("%f",&x)

	switch nomer{
	case 1:
		a=x;
        c=a* math.Sqrt(2);
        h=c/2;
        s=c*h/2;
        fmt.Println("gipotenuza  balandlik  Yuza ",c,h,s);
  
   case 2:
       c=x;
       a=c/math.Sqrt(2);
       h=c/2;
       s=c*h/2;
       fmt.Println("katet , balandlik,yuza",a,h,s);

   case 3:
       h=x;
       c=2*h;
       a=c/math.Sqrt(2);
       s=c*h/2;
       fmt.Println("katet,gipotenuza, yuza",a,c,s);

   case 4:
       s=x;
       h=math.Sqrt(2);
       c=2*h;
       a=math.Sqrt(2);
       fmt.Println("Katet,gipotenuza,balandlik",a,c,h);
	  default:
		fmt.Println("To'gri tanlov qiling") 

	}

}