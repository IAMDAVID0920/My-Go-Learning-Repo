package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	// get sum of 1+2+3+4+5
	var sum int = 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	scores := [5]int{33, 22, 11, 55, 66}
	for _, val := range scores {
		fmt.Printf("this is: %d\n", val)
	}

}
