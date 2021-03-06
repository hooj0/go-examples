// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-30
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
switch mulit value
==================================================================
switch 支持多个值一起选择判断，作为相同的条件值
------------------------------------------------------------------*/
package main

import "fmt"

func age(x int) {

	// 当 switch 包含参数，只能用固定值进行匹配
	// 没有条件参数，可以进行表达式判断
	switch x {
		case 1, 2, 3, 4, 5, 6:
			fmt.Println("幼儿园")
		case 7, 8, 9, 10:
			fmt.Println("小学")
		case 11:
			fmt.Println("初中")
		default:
			fmt.Println("高中")

	}
}

func main() {

	age(11)
	age(8)
	age(2)
	age(3)

	// 初中
	// 小学
	// 幼儿园
	// 幼儿园
}
