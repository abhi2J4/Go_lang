package main

import "fmt"

func main() {
	age := 45
	name := "Anu"
	height := 5.6777877

	fmt.Println("age",age,"name",name ,"height",height)

	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)
	fmt.Printf("Type is height %T\n", height)
	fmt.Printf("Type is age %T\n", age)
	//understand it
	fmt.Printf("Name:  %s ,Age: %d, Height: %.2f\n  ",name,age, height)

}
