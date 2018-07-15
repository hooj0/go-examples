// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-15
// copyright by hoojo@2018
// --------------------------------------------------------------




/* ---------------------------------------------------------------
func closures
==================================================================
func闭包 函数相当于一个函数工厂，每次调用都可以构建一个相同的func返回
------------------------------------------------------------------
* 学习函数闭包之前要先了解“函数类型参数”和“函数类型返回值”
------------------------------------------------------------------
func闭包 是一个函数值，是一个函数返回func的值
func闭包 函数闭包可以保护当前函数变量的作用域
func闭包 函数可以引用闭包函数之外的变量值，并且可以访问或设置变量值
func闭包 函数中的变量外部是不能直接访问和操作的
------------------------------------------------------------------*/
package main

import "fmt"

// 声明一个闭包函数adder()
// 返回值=func(int) int 返回一个函数
func adder() func(int) int {
	sum := 0

	return func(x int) int {
		// 访问外部的 sum 变量
		sum += x
		return sum
	}
}

func main() {

	// 调用 adder 方法，相当于在 adder方法中创建一个函数返回
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println("pos: ", pos(i), ", neg: ", neg(-2 * i))
	}

	// 闭包函数立即运行
	fmt.Println("run: ", adder()(3))
}