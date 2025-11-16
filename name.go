// package main
// import "fmt"
// func main(){
// 	fmt.Println("hello world")
// }

package main

import (
	"fmt"
	
	// "mylearning/data"
	// "mylearning/utils"
)

func main() {
	fmt.Println("hello jee")
	// fmt.Println("hello ghfdhjee")
	// fmt.Println(1 % 2)

	// utils.Printmassage("hello world")
	// data.Printmassage("hello je")

	//******* variable in main function  ******

	var name string = "jacky"
	var vesrion = "ekata dutta"

	fmt.Println(name)
		fmt.Println(vesrion)  //

 var money int = 900000
 var curreny = 79879
 
 fmt.Println(money)
 fmt.Println("currency : ", curreny)

 var dimesion float32 = 34.67
 fmt.Println(dimesion)

 var Answer bool = true
 Answer= false
 fmt.Println(Answer)

 const pi = 67.12
//  pi= 78
 fmt.Println(pi)


//  persion := "123"
 persion := "data "  // popular way to  every go developer used it
 fmt.Println(persion)

 var Public = "data is important "  // capturazation me aap anywhere used it and exoprt it
 var private = "data is not  important "

 fmt.Println(Public)
 fmt.Println(private)
}

