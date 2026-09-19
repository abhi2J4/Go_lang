// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sayHello() {
// 	fmt.Println("Hello, world!")
// 	// time.Sleep(2000 * time.Millisecond) // Simulating some work
// 	fmt.Println("sayHello function ended successfully")
// }

// func sayHi() {
// 	fmt.Println("Hi Prince :)")
// 	time.Sleep(1000 * time.Millisecond) // Simulating some work
// 	fmt.Println("Hi Prince Function ended:)")
// }

// func main() {
// 	fmt.Println("learning goroutines")

// 	go sayHello()
// 	go sayHi()

// 	// Wait for a moment to allow the goroutine to finish
// 	time.Sleep(800 * time.Millisecond)
// }

// **** If you mean print 1 to 10 in Go using a goroutine, here is the simple version: ***
package main

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println("print number 1 to 10 using gorutine")
	go printNumber(&wg)

	wg.Wait()

}


// *** Printing 1 to 10 using two goroutines




