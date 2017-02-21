// go routines - lightweight threads! (user threads)

package main


import "fmt"
import "time"

func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)

		// Pause the function for 1 second to allow other functions to execute
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {

	for i := 0; i < 10; i++ {
		go count(i)
	}

	time.Sleep(time.Millisecond * 11000)
}
