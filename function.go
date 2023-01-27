package main

import "fmt"

/*
func function_name(){
	statement-1
	statement-2
	statement-3
	statement-4
}
*/

// func main() {
// 	a()
// }
// func a() {
// 	fmt.Print(12)
// }
// func b() {
// 	c()
// }
// func c() {
// } //Output is 12 here.

// func main() {
// 	result := a()
// 	fmt.Println(result)
// }
// func a() int {
// 	return 122
// }
// func b() {
// 	c()
// }
// func c() {

// } //Output is 122

// func main() {
// 	result, x := a()
// 	fmt.Println(result)
// 	fmt.Println(x)
// }
// func a() (int, string) {
// 	return 122, "weqrr3"
// }
// func b() {
// 	c()
// }
// func c() {

// } //Output: 122 weqrr3

// func main() {
// 	result, x := c()
// 	fmt.Println(result)
// 	fmt.Println(x)
// }
// func a() (int, string) {
// 	return 122, "weqrr3"
// }
// func b() {
// 	c()
// }
// func c() (i int, j string){
// 	i=10
// 	j="rahul"
// 	return
// } //Output: 10 rahul

// func main() {
// 	result, x := c()
// 	fmt.Println(result)
// 	fmt.Println(x)
// 	b(1, 2, 3, 4, 5, 6, 6)
// }
// func a() (int, string) {
// 	return 122, "weqrr3"
// }
// func b(args	...int) {
// 	fmt.Println(args)
// }
// func c() (i int, j string){
// 	i=10
// 	j="rahul"
// 	return
// } //Output: 10 rahul [1 2 3 4 5 6 6]

// func main() {
// 	result, x := c()
// 	fmt.Println(result)
// 	fmt.Println(x)
// 	r := b(1, 2, 3, 4, 5, 6, 6)
// 	fmt.Print(r)
// }
// func a() (int, string) {
// 	return 122, "weqrr3"
// }
// func b(args	...int) bool {
// 	for _, v := range args {
// 		fmt.Print(v)
// 	}
// 	return true
// }
// func c() (i int, j string){
// 	i=10
// 	j="rahul"
// 	return
// } //Output: 10 rahul 1234566true

// func main() {
// 	result, x := c()
// 	fmt.Println(result)
// 	fmt.Println(x)
// 	r, _ := b(1, 2, 3, 4, 5, 6, 6)
// 	fmt.Print(r)
// }
// func a() (int, string) {
// 	return 122, "weqrr3"
// }
// func b(args ...int) (bool, int) {
// 	for _, v := range args {
// 		fmt.Println(v)
// 	}
// 	return true, 11
// }
// func c() (i int, j string) {
// 	i = 10
// 	j = "rahul"
// 	return
// }

func main() {
	w := new(int)
	name := new(string)
	result, x := c(w, name)
	fmt.Println(result)
	fmt.Println(x)
	fmt.Println(*name)
	r, _ := b(1, 2, 3, 4, 5, 6, 6)
	fmt.Print(r)
}
func a() (int, string) {
	return 122, "weqrr3"
}
func b(args ...int) (bool, int) {
	for _, v := range args {
		fmt.Println(v)
	}
	return true, 11
}
func c(w *int, name *string) (i int, j string) {
	i = 10
	j = "abhay"
	*w = 100
	*name = "code"
	return
}
