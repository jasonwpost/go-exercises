package main

import "fmt"

func main() {
  yourAge := 18

  if yourAge >= 16 {
    fmt.Println("You Can Drive")
  } else if yourAge >= 18 {
    fmt.Println("You Can Vote")
  } else {
    fmt.Println("You Can Have Fun")
  }

  yourAge = 5

  switch yourAge {
    case 16: fmt.Println("Go Drive")
    case 18: fmt.Println("Go Vote")
    default:fmt.Println("Go Have fun")
  }

}
