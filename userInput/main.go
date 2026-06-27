package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("hello ,what is your name?")

	// var name string;
	

	

	// fmt.Scan(&name2)
	// fmt.Println("my age is ",name2)
	// fmt.Scan(&name)
	

	reader := bufio.NewReader(os.Stdin)
	name, _:= reader.ReadString('\n')
    fmt.Println("my name is :",name)



	
}