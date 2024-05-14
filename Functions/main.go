package main

import "fmt"

func main() {
	fact := factorial((5))
	fmt.Println(fact)

}




// func factorial(number int) int {
// 	result := 1
// 	for i := 1; i <= number; i++ {
// 		result = result * i
// 	}
// 	return result
// }

//factorial of 5 => 5*4*3*2*1 => 120
