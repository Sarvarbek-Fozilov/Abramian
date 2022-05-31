package main

import (
	"fmt"
	"math"
)
 
func main(){
	var nomer int
	var x,radius,d,l,s float64

	fmt.Println("Elementlarni tanlang\n 1.radius 2.diametr 3.uzunlik 4.doira yuzi")
	fmt.Scanf("%d",&nomer)
	fmt.Println("Doiraning radiusini kirit")
	fmt.Scanf("%f",&x)

	switch nomer{
	case 1:
		radius=x
		d=2*radius
		l=2*3.13*radius
		s=3.14*radius*radius
		fmt.Println(d,l,s)
	case 2:
        radius=x/2;
        d=x;
        l=2*3.14*radius;
        s=3.14*radius*radius;
        fmt.Println(radius,l,s);
        break;
   case 3:
        radius=x/2*3.14;
        d=2*radius;
        l=x;
        s=3.14*radius*radius;
        fmt.Println(radius,d,s);
        break;
   case 4:
        radius=math.Sqrt(x/3.14)
        d=2*radius;
        l=2*3.14*radius;
        s=x;
        fmt.Println(radius,d,l);
	}

}