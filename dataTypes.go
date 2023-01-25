package main

import "fmt"

func main() {
	//var Variable_name datatype
	var data int
	data = 10

	data1 := 100 // declare(automatic) the variable and assign value to it
	//data :=1000  //redeclare the data variable
	data = 1000
	Data := 1000
	//data1 = false  //wrong assignment value type
	fmt.Println(data)
	fmt.Println(data1)
	fmt.Println(Data)
	//data type
	//1. Primitave
	//int , float64 , string, bool, complex
	//2. non- primitave
	//struct, map, chan, interface, array, slice
	//functional programming
	//recent feature update of java - from java 11
	name := "Abhay"
	fmt.Println(name)
}
