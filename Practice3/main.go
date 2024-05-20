package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	greeting := "Hello"
	for i := 0; i < 5; i++ {
		fmt.Println(greeting, i)
	} 
}