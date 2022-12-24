package main

import (
	"fmt"
	"reflect"
)

// slice
// 其实是个结构体

func main() {
	slice := make([]int, 4, 20)
	slice[0] = 66
	slice[1] = 88
	slice[2] = 99
	slice[3] = 100

	// loop 2 ways
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v \n", i, slice[i])
	}

	fmt.Println("----------------------------")

	// for range loop
	for i, v := range slice {
		fmt.Printf("index: %v, element: %v \n", i, v)
	}

	// slice 不能直接使用 需要先引用到一个数组 要么 make slice去引用
	// var slice1 []int
	// fmt.Println(slice1, reflect.TypeOf(slice1))
	slice1 := make([]int, 4, 20)
	fmt.Println(slice1, reflect.TypeOf(slice1))

	var intarr [6]int = [6]int{1, 4, 7, 2, 5, 8}
	var slice2 []int = intarr[1:4]
	fmt.Println(slice2[0])
	fmt.Println(slice2[1])
	fmt.Println(slice2[2])
	// // will out of bound of this slice [1:4] range!
	// fmt.Println(slice2[3])

	// You could always slice a slice

	// slice len can increase dynamically
	// use append
	// 底层原理
	// 1.追加元素 实际是对数组进行enlarge, 老数组enlarge成新数组
	// 2.创建一个新的数组, 把老的复制过去 新数组再追加
	slice3 := append(slice2, 70, 23)
	fmt.Println(slice3)

	// copy slice
  var a []int = []int{1, 4, 7, 2, 5, 8}
  var b []int = make([]int, 10)
  // copy
  copy(b, a) // 将a中对应元素复制到b中对应里
  fmt.Println(b)
}
