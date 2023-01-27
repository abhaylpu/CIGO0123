package main

import "fmt"

func main() {
	// 	for i:=0;i<3;i++ {
	// 		for i:=0;i<3;i++ {
	// 			fmt.Println(i)
	// 		}
	// 	}

	// i := 0
	// for i<3 {
	// 	if i==1 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	break
	// }

	// for i := 0; i<3; i++ {
	// 	if i == 1 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	//break
	// }
	// list:=range 2

	// nums:=[]int{1, 3, 2, 4, 0}
	// for k, v:=range nums {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }

	// nums:=[]int{1, 3, 2, 4, 0}
	// for index, value := range nums {
	// 	fmt.Println(index)
	// 	fmt.Println(value)
	// }

	// nums:=[]int{1, 3, 2, 4, 0}
	// for _, value := range nums {
	// 	fmt.Println(value)
	// }

	// nums:=[]int{1, 3, 2, 4, 0}
	// for _, value := range nums {
	// 	if value == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(value)
	// }

	// nums:=[]int{1, 3, 2, 4, 0}
	// for _, value := range nums {
	// 	if value == 3 {
	// 		break
	// 	}
	// 	fmt.Println(value)
	// }

	// for _, value:=range "rahul" {
	// 	fmt.Println(value)
	// }  //Here it prints the ASCII values of the word "rahul"

	nums := []string{"abhay", "pratap", "singh"}
	for _, value := range nums {
		fmt.Println(value)
	}
}
