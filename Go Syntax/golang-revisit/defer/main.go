package main

import "fmt"

func main() {
	test()
	fmt.Println("before test been executed successfully")
	fmt.Println("hello")

}

// defer+recover instead of try/exception
// using recover to get panic

func test() {

	defer func() {
		// 不会让程序轻易error out

		err := recover()
		// if not catch error, return nil
		if err != nil {
			// catch the error
			fmt.Println("error has been catched, error is: ", err)
		}
	}() //defer 后加入匿名函数调用
	num1 := 10
	num2 := 0

	res := num1 / num2
	fmt.Print(res)
}
