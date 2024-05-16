package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 5, 4, 5, 6}
	sum := sumup(numbers)
	fmt.Println(sum)
	
	sumVariatic := sumupVariatic(1, 2, 3)
	fmt.Println(sumVariatic)

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