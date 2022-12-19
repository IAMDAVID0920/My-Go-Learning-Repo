package main

import (
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	pl(os.Args)
	args := os.Args[1:]
	pl(args)
	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}

	fmt.Println("Max Value: ", max)

}
