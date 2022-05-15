//To'gri uchburchakning katetlari a va b berilgan .
//Uning gipetonuzasi c va perimertrianiqlansin c=sqrt(a^2+b^2)p = a+b+c

package main
import (
  "fmt"
  "math"
)
func main(){
  var a,b,p float64
  fmt.Println("Birinchi tomondi kiriting")
  fmt.Scanf("\n%f", &a)

  fmt.Println("Ikkinchi tomondi kiriting")
  fmt.Scanf("\n%f", &b)


  res := math.Sqrt   ((a*a)+(b*b))
  p= a+b+res

  fmt.Println("Gipotenuza:",res)
  fmt.Println("Perimetr:",p)
}