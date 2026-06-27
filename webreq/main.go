package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learnig web Service  ")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Print("error getting response ", err)
		return
	}

	defer res.Body.Close() // resource manage is important in goLang
	fmt.Printf("Type of response : %T\n", res)
	//    fmt.Println(res)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response", err)
	}
	fmt.Println("data: ", string(data))

}
