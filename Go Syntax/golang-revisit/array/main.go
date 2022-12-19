package main

import (
	"fmt"
)

func main() {
	// Array initialization
	var arr1 [3]int = [3]int{3, 6, 9}
	fmt.Println(arr1)

	var arr2 = [3]int{1, 4, 7}
	fmt.Println(arr2)

	var arr3 = [...]int{4, 5, 6, 7}
	fmt.Println(arr3)

	var arr4 = [...]int{2: 66, 0: 33, 1: 99, 3: 88}
	fmt.Println(arr4)

	// array1 type is: [3]int
	// array4 type is: [4]int
	// 值拷贝
	fmt.Printf("array1 type is: %T\n", arr1)
	fmt.Printf("array4 type is: %T\n", arr4)

	var testArr = [3]int{3, 6, 9}
	test1(&testArr)
	fmt.Println(testArr)

	// 2d array
	var arr [2][3]int16
	// default [[0 0 0] [0 0 0]]
	fmt.Println(arr)
	fmt.Printf("arr address: %p\n", &arr)
	fmt.Printf("arr[0] address: %p\n", &arr[0])
	fmt.Printf("arr[0][0] address: %p\n", &arr[0][0])
	fmt.Printf("arr[1] address: %p\n", &arr[1])
	fmt.Printf("arr[1][0] address: %p\n", &arr[1][0])

	// give value
	arr[0][0] = 82
	arr[0][1] = 47
	arr[1][1] = 25

	fmt.Println(arr)

	// init 2d arr
	var arr2D [2][3]int = [2][3]int{{1, 3, 7}, {2, 5, 8}}
	fmt.Println(arr2D)

	// iterate 2d arr
	// 1 way
	var arr2Dver2 [3][3]int = [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	for i := 0; i < len(arr2Dver2); i++ {
		for j := 0; j < len(arr2Dver2[i]); j++ {
			fmt.Print(arr2Dver2[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println("---------------------")
	// another way

	for key, val := range arr2Dver2 {
		// val will be the nested array
		for k, v := range val {
			fmt.Printf("arr[%v][%v] = %v\t", key, k, v)
		}
		fmt.Println()
	}
}

func test1(arr *[3]int) {
	(*arr)[0] = 7
}
