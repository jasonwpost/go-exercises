package main

import "fmt"

// var x int // package level scope // var sets the variable to 0

func main(){
  x := 7 // short variable declaration operator // block level scope
  fmt.Printf("%T", x) // prints the type

  // type {} // composite literal

  xi := []int{2,4,7,9,42} // slice holds a list of things ~ Array
  fmt.Println(xi)

  m := map[string]int{ // map holds key value pairs ~ Hash Table
    "Todd": 45,
    "Job": 42, // nessecary to have the trailing comma for good code practice
  }
  fmt.Println(m)

  p1 := person{ // if values are given in order then you do not have to
    "Miss", // provide the field name
    "Moneypenny",
  }
  fmt.Println(p1)

  p1.speak()

  sa1 := secretAgent{
    person{
      "James",
      "Bond",
    },
    true,
  }

  sa1.speak()

  // a problem now - a secret agent has two potential speak methods
  // but we can access the deeper type through specifying
  sa1.person.speak()
  sa1.speak()

  saySomething(p1)
  saySomething(sa1)

  // this is an anon struct
  p2 := struct {
    fname string
    lname string
  }{
    "James",
    "Bond",
}

}

type person struct { // struct is ~ object
  fname string // lowercase f means not visible outside the package
  lname string // Uppercase would mean it is visible
}

// can also attach function to a struct
func (p person) speak() {
  fmt.Println(p.fname, p.lname, `says, "Good Morning, Jason."`)
}

// object composition
type secretAgent struct {
  person
  licenseToKill bool
}

func (sa secretAgent) speak() {
  fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

// interfaces
// this is a little weird, but because person and secretAgent both have speak methods,
// they are also automatically of type human, without any extra declaration
type human interface {
  speak()
}

// the above is then demonstrated by this method
func saySomething(h human){
  h.speak()
}
