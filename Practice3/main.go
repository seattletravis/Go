package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var greeting string = "Hello"


	for i := 0; i < 5; i++ {
		fmt.Println(greeting, i)
	} 
}