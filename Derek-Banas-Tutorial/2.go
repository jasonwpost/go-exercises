package main

import "fmt"

func main() {
  const pi float64 = 3.14159265

  /*var (
    varA = 2
    varB = 3

  )*/

  var myName string = "Jason"

  fmt.Println(len(myName))
  fmt.Println(myName + " is a robot")
  fmt.Println("I like \n")
  fmt.Println("I like \n")

  var isOver40 bool = true

  fmt.Printf("%f \n", pi)
  fmt.Printf("%T \n", pi)
  fmt.Printf("%T \n", isOver40)
  fmt.Printf("%d \n", 100)
  fmt.Printf("%b \n", 100)
  fmt.Printf("%c \n", 44)
  fmt.Printf("%x \n", 100)
  fmt.Printf("%e \n", pi)

  fmt.Println("true && false =" , true && false)
  fmt.Println("true || false =" , true || false)
  fmt.Println("!true =" , !true)


}
