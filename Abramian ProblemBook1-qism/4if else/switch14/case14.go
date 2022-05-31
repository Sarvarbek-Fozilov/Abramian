package main

import (
	"fmt"
	"math"
)

func main(){
	var nomer int
	var x,a,r1,r2,s float64

	fmt.Println("Elementni tanlang\n 1.Tomon 2.ichki chizilgan aylana radiusi 3.tashqi chizilgan aylana radiusi 4 yuzasi")
	fmt.Scanf("%d",&nomer)
	fmt.Println("Uzunlikni kiriting")
	fmt.Scanf("%f",&x)

	switch nomer{
	case 1:

		a=x;
        r1=a*(math.Sqrt(3))/6
        r2=2*r1;
        s=a*a*(math.Sqrt(3))/4;
        fmt.Println(" Ichki aylana Radiusi,tashqi aylana radiusi , Yuzasi",r1,r2,s);

   case 2:
       r1=x;
       a=r1*6/(math.Sqrt(3));
       r2=2*r1;
       s=a*a*(math.Sqrt(3))/4;
       fmt.Println("Katet,Tashqi aylana radiusi , yuzasi",a,r2,s);

   case 3:
       r2=x;
       r1=r2/2;
       a=r1*6/(math.Sqrt(3));
       s=a*a*(math.Sqrt(3))/4;
       fmt.Println("Katet,Ichki ayalana radiusi, yuzasi",a,r1,s);
  
   case 4:
       s=x;
       a=math.Sqrt(s*4/(math.Sqrt(3)));
       r1=a*(math.Sqrt(3))/6;
       r2=2*r1;
       fmt.Println("Katet,Ichki aylana radiusi,Tashqi aylana radiusi",a,r1,r2);
   default:
	fmt.Println("To'g'ri  birlik tanlang")
        
	}
}