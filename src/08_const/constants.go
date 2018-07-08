// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------




/* ---------------------------------------------------------------
const
==================================================================
常量 用于声明一个固定不变，全局可用的变量数据
------------------------------------------------------------------
const PI = 3.14159267
------------------------------------------------------------------
常量 声明与变量类似，只不过是使用 const 关键字
常量 可用被外部访问
常量 可以是字符、字符串、布尔值或数值类型
常量 不能用 := 语法声明。
------------------------------------------------------------------*/


package main

import "fmt"

// 常量声明
const Pi  = 3.14159267
const MAX_VALUE  = 10122111.2

// 常量组
const (
	a = 1
	b = true
	c = "str"
)

func main() {

	fmt.Printf("Pi: %v, MAX_VALUE: %v \n", Pi, MAX_VALUE)
	fmt.Printf("a: %v, b: %v, c: %v \n", a, b, c)

	// 常量声明
	const Valid  = false
	fmt.Println("Valid: ", Valid)

}
