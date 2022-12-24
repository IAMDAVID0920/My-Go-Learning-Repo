package main

import (
	"fmt"
)

func main() {
	// var 变量名 map[keytype]valuetype
	var thisMap map[int]string
	fmt.Println(thisMap)
	fmt.Println("-----------")
	// thisMap[0] = 20
	// thisMap[1] = 21
	// fmt.Println(thisMap)
	thisMap = make(map[int]string, 10) //map could save 10 key-value pairs

	thisMap[200095452] = "张三"
	thisMap[200095387] = "李四"
	thisMap[200097291] = "王五"
	fmt.Println(thisMap)
	// map must use make to initialize
	// key cannot duplicate, if duplicate the latter will replace first one

	// way 2
	b := make(map[int]string)
	fmt.Println(b)
	b[200095452] = "张三"
	b[200095387] = "李四"
	b[200097291] = "王五"
	fmt.Println(b)

	// way 3
	c := map[int]string{
		200095387: "李四",
		200097291: "王五",
	}
	c[200095452] = "张三"
	fmt.Println(c)

	// delete(map, key)
	// delete(b, 200095452)
	// fmt.Println(b)

	// search 查找
	// value, bool = map[key] bool: either true or false
	val, flag := b[200095452]
	fmt.Println(val, flag)

	//get len
	fmt.Println(len(b))

	// map only accept -> for range to do the iteration
	for k, v := range b {
		fmt.Printf("%v:%v", k, v)
	}

	a := make(map[string]map[int]string)
	// key is string, val is a key int val str map
	a["班级1"] = make(map[int]string, 3)
	a["班级1"][20096677] = "lulu"
	a["班级1"][20096677] = "lili"
	a["班级1"][20096677] = "feifei"

	a["班级2"] = make(map[int]string, 3)
	a["班级2"][20096677] = "leilei"
	a["班级2"][20096677] = "aa"
	a["班级2"][20096677] = "kk"

	for k1, v1 := range a {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("student number: %v, student name: %v", k2, v2)
		}
		fmt.Println("----------")
	}
}
