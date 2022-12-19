package main

import (
	"fmt"
)

func main() {
	// fmt.Println("3 + 5 = ", sum(3, 5))
	sum1, diff1 := Cal(3, 5)
	// if you just want to ignore one result, use underscore _
	sum2, _ := Cal(7, 2)
	fmt.Printf("sum1 = %v, diff1 = %v, sum2 = %v\n", sum1, diff1, sum2)

	test()
	fmt.Println("------------")
	test(37, 25)

	var num int = 10
	test2(num)
	fmt.Println("main ----- ", num)

	fmt.Println(&num)
	test3(&num) // call test3, passin num address
	fmt.Println("main --- ", num)

	// function could be a data type in Go, could be assigned to a variable
	a := sum
	fmt.Printf("type of a is: %T, type of sum is: %T \n", a, sum)
	// res := a(10, 20)
	// a(10, 20) <-> sum(10, 20) exactly the same
	// fmt.Println(res)

	test02(10, 3.14, sum)
	test02(10, 3.14, a)
}

func test02(num1 int, num2 float64, testFunc func(int, int)) {
	fmt.Println("-----test02")
}

// 大写相当于public 小写相当于private
func sum(a int, b int) {
	fmt.Println(a + b)
}

func Sum(a int, b int) int {
	return a + b
}

func Cal(a int, b int) (int, int) {
	return a + b, a - b
}

// declare a func using ... as 可变参数
func test(args ...int) {
	// can passin any number of int value 0 to n
	// how to deal with it
	// 将可变参数变成切片
	// iterate

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

// 默认value passin/copy 值传递, what you've modified in function won't mess up with original value
// end of the func, will disappear from the stack

func test2(num int) {
	num = 30
	fmt.Println("test --- ", num)
}

// if you wish to change variable value inside the function from outside passed in
// pass in using variable address &

func test3(num *int) {
	// arg type is int pointer
	// need to pass in int address

	*num = 30
	fmt.Println("test3 ----", *num)
	// one * corresponding with one &
}
