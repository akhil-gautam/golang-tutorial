package main

import "fmt"

// func main(){
// 	var arr [3][2]int
// 	arr[0][0] = 1
// 	arr[0][1] = 2
// 	arr[1][0] = 3
// 	arr[1][1] = 4
// 	for i := 0; i < len(arr); i++{
// 		for j := 0; j < len(arr[0]); j++{
// 			fmt.Printf("%v ",arr[i][j])
// 		}
// 		fmt.Printf("\n")
// 	}
// }

func main(){
	var length int
	fmt.Println("Enter the size of array: ")
	fmt.Scanln(&length)

	var arr = make([]int, length)

	for i := 0; i < length; i++{
		fmt.Printf("Enter %v number\t", i+1)
		fmt.Scanln(&arr[i])
	}
	for i := 0; i < length; i++{
		fmt.Println(arr[i])
	}
}