// structs

package main

import "fmt"

func main() {

  rect1 := Rectangle{leftX: 0, topY:50, height:10,width:10}
  // is the same as
  rect2 := Rectangle{0, 50, 10, 10}

  fmt.Println("Rectangle 1 is", rect1.width,"wide")
  fmt.Println("Rectangle 2 is", rect2.width,"wide")

  fmt.Println("Area of Rect =", rect1.area())

}

type Rectangle struct {
  leftX float64
  topY float64
  height float64
  width float64
}

func (rect *Rectangle) area() float64{
  return rect.width * rect.height
}
