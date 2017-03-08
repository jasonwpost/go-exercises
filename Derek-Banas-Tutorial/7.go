package main

import "fmt"

func main() {
  listNums := []float64{1, 2, 3, 4, 5}

  fmt.Println("sum:", addThemUp(listNums))

  num1, num2 := next2Values(5)
  fmt.Println(num1, num2)

  fmt.Println(subtractThem(1,2,3,4,5))

  // closures - allows for shared variables
  // between functions

  num3 := 3

  doubleNum := func() int {
    num3 *= 2
    return num3
  }

  fmt.Println(doubleNum())

  fmt.Println(factorial(3))

  // defer executes function after main function ends
  defer printTwo()
  printOne()

  fmt.Println(safeDiv(3, 0))
  fmt.Println(safeDiv(3, 2))

  demPanic()

}

// (arguments) return type
func addThemUp(numbers []float64) float64 {
  sum := 0.0

  for _, val := range numbers {
    sum += val
  }
  return sum
}

// two return values!
func next2Values(number int) (int, int){
  return number+1, number+2
}

func subtractThem(args...int) int {
  finalValue := 0

  for _, value := range args {
    finalValue -= value
  }

  return finalValue
}

func factorial(num int) int {
  if num == 0 {
    return 1
  }
  return num * factorial(num - 1)
}

func printOne() { fmt.Println(1)}
func printTwo() { fmt.Println(2)}

// If an error occurs we can catch the error with recover and allow
// code to continue to execute
func safeDiv(num1, num2 int) int {
  defer func(){
    fmt.Println(recover())
  }()

  solution := num1 / num2
  return solution
}

func demPanic(){
  defer func(){
    fmt.Println(recover())
  }()

  panic("PANIC")

}
