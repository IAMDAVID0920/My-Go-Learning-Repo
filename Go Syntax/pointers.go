package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Print("hello, world\n")

	i, j := 42, 2701
	fmt.Println(i, j)
	// print the address of i and address of j
	fmt.Println(&i, &j, reflect.TypeOf(&i), reflect.TypeOf(&j))
	p := &i
	fmt.Println(*p, p, reflect.TypeOf(*p), reflect.TypeOf(p))

	*p = 21
	fmt.Println(i, &i)

	arr := make([]int, 5)
	fmt.Print(arr)

}
