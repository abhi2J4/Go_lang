package main

import "fmt"

func simpleFunction(){
	fmt.Println("This is my function")
}

func add(a , b int ) int{
	return  a + b
}

func multiply(a , b int ) (result int ){
	return  a * b
	return
}

func main() { 
	fmt.Println("we are learning in goLang")
	simpleFunction()
	ans := add(7,9)
	
	fmt.Println(ans)
	data := multiply(3,4)
	fmt.Print(data)
}