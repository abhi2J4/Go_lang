// In Go, we use a channel to allow goroutines to communicate and safely exchange data.

// Online Go compiler to run Golang program online
// Print "Start small. Ship something." message

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func reader(id int, ch <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for {
// 		val, ok := <-ch
// 		if !ok {
// 			fmt.Println("channel is closed")
// 			return
// 		}
// 		fmt.Printf("Reader %d Received %d\n", id, val)
// 	}

// }

// func main() {

// 	ch := make(chan int)

// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go reader(1, ch, &wg)
// 	go reader(2, ch, &wg)
// 	go reader(3, ch, &wg)
// 	go reader(4, ch, &wg)

// 	for i := 0; i < 100; i++ {
// 		ch <- i
// 	}

// 	close(ch)
// 	wg.Wait()

// }

// Goroutine 1
//     |
//     |  send data
//     ↓
//  Channel
//     |
//     |  receive data
//     ↓
// Goroutine 2


// Show me how to print 1–10 using two goroutines
package main

import (
	"fmt"
	"sync"
)

func odd(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i += 2 {
		<-oddCh

		fmt.Println(i)

		evenCh <- true
	}
}

func even(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= 10; i += 2 {
		<-evenCh

		fmt.Println(i)

		if i < 10 {
			oddCh <- true
		}
	}
}

func main() {
	oddCh := make(chan bool)
	evenCh := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go odd(oddCh, evenCh, &wg)
	go even(oddCh, evenCh, &wg)

	// Start with odd number 1
	oddCh <- true

	wg.Wait()
}