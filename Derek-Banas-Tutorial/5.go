package main

import "fmt"

func main() {
  var favNums2[5] float64

  favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618

  fmt.Println(favNums2[3])

  favNums3 := [5]float64 {1, 2, 3, 4, 5}

  for i, value := range favNums3 {
    fmt.Println(value, i)
  }
  for _, value := range favNums3 {
    fmt.Println(value)
  }

  numSlice := []int { 5, 4, 3, 2, 1 }

  numSlice2 := numSlice[3:5]

  fmt.Println("numSlice2[0] =", numSlice2[0])

  // You can also create an empty slice and define the data type,
	// length (receive value of 0), capacity (max size)

  numSlice3 := make([]int, 5, 10)

  copy(numSlice3, numSlice)

  fmt.Println(numSlice3[0])

  numSlice3 = append(numSlice3, 0, -1)

	fmt.Println(numSlice3[6])

}
