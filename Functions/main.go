package main

import "fmt"

func main() {

	// call func with slice
	numbers := []int{1, 2, 3, 5, 4, 5, 6}
	sum := sumup(numbers)
	fmt.Println(sum)
	
	// call variatic func with multiple numbers
	sumVariatic := sumupVariatic(1, 2, 3, 4, 5)
	fmt.Println(sumVariatic)

	// call variatif func with slice
	sumCallVariaticWithSlice := sumupVariatic(numbers...)
	fmt.Println(sumCallVariaticWithSlice)

}

// function that takes array of ints
func sumup(numbers []int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

// variatic version of function
func sumupVariatic(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}