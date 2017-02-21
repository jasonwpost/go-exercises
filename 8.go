package main

import "fmt"

func main() {

  x := 0

  changeXVal(x)

  fmt.Println("x =", x)

  changeXValNow(&x)

  fmt.Println("x =", x)

  fmt.Println("memory address of x =", &x)

  yPtr := new(int)
  changeYValNow(yPtr)

  fmt.Println("y =", *yPtr)

}


// false example - won't change x
func changeXVal(x int){
  x = 2
}

// true example - will change x
func changeXValNow(x *int){
  *x = 2
}

// true example - will change y
func changeYValNow(yPtr *int){
  *yPtr = 100
}
