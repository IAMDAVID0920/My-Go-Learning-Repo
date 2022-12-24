package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	var t *Teacher = new(Teacher)
	// t is pointer, which means t is address
	// 给地址指向的对象的字段赋值
	(*t).Name = "dd"
	(*t).Age = 12
	//go底层弄完了
	(*t).School = "hku"
	t.School = "osu"

	fmt.Println(*t)

	/////////////////

	var t1 *Teacher = &Teacher{}
	(*t1).Name = "dd"
	(*t1).Age = 12
	//go底层弄完了
	(*t1).School = "hku"
	t1.School = "osu"
	fmt.Println(*t1)

}

// func main() {
// 	var t Teacher
// 	fmt.Println(t)
// 	t = Teacher{"david", 12, "ohio state university"}
// 	fmt.Println(t)
// }
