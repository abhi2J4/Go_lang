package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Print("Error will be craeted  file ", err)
		}
		defer file.Close()

		// second part
		content := "hello world  by Abhishek "
		byte, errors := io.WriteString(file, content)
		fmt.Print("byte written is : ", byte)
		if errors != nil {
			fmt.Println("Error will be writing the file ", errors)
			return
		}
	*/

	//Reading the file

	/*
		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Print("Error WHile  open  file ", err)
			return
		}
		defer file.Close()

		//creting the bufer to  read the file content

		buffer := make([]byte, 1024)

		//Read the file content  into the buffer
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while is Reading file", err)
				return
			}

			//Process the read the content
			fmt.Println(string(buffer[:n]))


		}
	*/

	//Read the Entire  file into  a byte slice

	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading the file", err)
		return
	}
	fmt.Println(string(content))
}

//buffer is tempory kind of memory  that store the data
