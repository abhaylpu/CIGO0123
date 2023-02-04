package main

import "fmt"

func main() {
	// i := 10
	// j := &i
	// *j = 100
	// fmt.Println(i)
	// fmt.Println(j) //Output: 100 0xc000018098

	// var i int
	// i = 10
	// j := new(int)
	// j = &i
	// *j = 100
	// fmt.Println(*j)
	// fmt.Println(i) //Output: 100 100

	// var i int
	// i = 10
	// j := new(int)
	// fmt.Println(j)
	// fmt.Println(i) //Output: 0xc000018098 10

	// var i int
	// i = 10
	// j := new(int)
	// j = &i
	// *j = 100
	// name := new(string)
	// fmt.Println(name) //Output: 0xc000052260

	var i int
	i = 10
	j := new(int)
	j = &i
	*j = 100
	name := new(string)
	*name = "golang"
	fmt.Println(*name) //Output: golang
}
