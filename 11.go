package main

import ("fmt"
"strings"
"sort"
"os"
"log"
"io/ioutil"
"strconv")


func main() {

  //strings
  sampString := "Hello World"

  // true if sampString contains 'lo'
  fmt.Println(strings.Contains(sampString, "lo"))
  // other functions
  fmt.Println(strings.Index(sampString, "lo"))
  fmt.Println(strings.Count(sampString, "l"))
  fmt.Println(strings.Replace(sampString, "l", "x", 3))

  csvString := "1,2,3,4,5,6"
  fmt.Println(strings.Split(csvString, ","))

  listOfLetters :=[]string{"c", "a", "b"}
  sort.Strings(listOfLetters)
  fmt.Println("Letters:", listOfLetters)

  // returns a string using the values passed in seperated
  listOfNums := strings.Join([]string{"3", "2", "1"}, ", ");

  fmt.Println(listOfNums);

  // File I/0

  file, err := os.Create("samp.txt")

  if err != nil {
    log.Fatal(err)
  }

  file.WriteString("This is some random text")

  file.Close()

  stream, err := ioutil.ReadFile("samp.txt")

  if err != nil {
    log.Fatal(err)
  }

  // convert into string
  readString := string(stream)

	fmt.Println(readString)

  // accepting input

  fmt.Println("What is your name? ")

  var name string

  fmt.Scan(&name)

  fmt.Println("Hello", name)

  //  casting

	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	// Convert number types
	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))

	// Convert a string into an int
	newInt, _ := strconv.ParseInt(randString, 0, 64)
    fmt.Println(newInt)

    // Convert a string into a float
    newFloat, _ := strconv.ParseFloat(randString2, 64)
    fmt.Println(newFloat)

}
