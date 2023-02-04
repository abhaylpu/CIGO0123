package main

import "fmt"

func main() {
	// number := 10
	var number int
	number = 10
	fmt.Println(number)

	//store a function as a value to a variable.
	var getInt func(int) int
	getInt = func(x int) int {
		fmt.Println("In 1st function")
		return 20 + x
	}
	g(getInt)
	// getInt(1)
	// getInt = func(x int) int {
	// 	fmt.Println("In 2nd function")
	// 	return 10 + x
	// }
	var y func()
	y = func() {
		fmt.Print(9)
	}
	y()
	j := func(x int) int {
		fmt.Println("In a 1st function")
		return 20 + x
	}(10)
	fmt.Println(j)
}
func g(getInt func(int) int) {
	getInt(78)
}

//function is a first class member in golang
