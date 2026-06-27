package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Start learning defer ")
	data := add( 2,7)
	fmt.Println("add two number is :",data)
	defer fmt.Println("middle learning  ")
	defer fmt.Println("middle learning new   ")
	fmt.Println("Last Learning ")

}
// A stack is LIFO (Last In, First Out) working 

// Start learning defer 
// add two number is : 9
// Last Learning 
// middle learning new   
// middle learning
