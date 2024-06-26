package main

import f "fmt"

type Product struct {
	title string
	id string
	price float64
}

func main () {
	// 1
	hobbies := [3]string{"landscaping", "cooking", "woodworking"}
	f.Println(hobbies)
	// 2
	f.Println(hobbies[0])
	f.Println(hobbies[1:])
	// 3
	mainHobbies := hobbies[:2]
	f.Println(mainHobbies)

	f.Println(cap(mainHobbies))
 	// 4
	mainHobbies = mainHobbies[1:3]
	f.Println(mainHobbies)
	// 5
	goals := []string{"learn GO,", "master GO,"}
	// 6
	goals[1] = "Become a GO God!,"
	goals = append(goals, "get paid,")
	f.Println(goals)
	// 7
	mine := []Product{
		{
			"foo", 
			"1", 
			2.99,
		}, 
		{
			"bar", 
			"2", 
			3.99,
		},
		}	
	f.Println(mine)

	mine = append(mine, Product{"buzz", "3", 89.99})
	f.Println(mine)


}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.