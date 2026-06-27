package main

import "fmt"

func divide(a, b int) (int ,error){
if b == 0{
	return 0,fmt.Errorf("error")
}
return a / b, nil
}

func main() {
	fmt.Println("Started the error handling")
	ans, _ := divide(10,5)
	fmt.Println(ans)
}