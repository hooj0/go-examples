// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
for
==================================================================
循环语句 用来迭代、遍历或重复做相同的一段业务，Go 只有一种循环结构：for 循环。
------------------------------------------------------------------
基本的 for 循环由三部分组成，它们用分号隔开：

	初始化语句：	在第一次迭代前执行
	条件表达式：	在每次迭代前求值
	后置语句：	在每次迭代的结尾执行

最后的大括号是必须的，不能省略，即便里面没有内容

语法：
	for 初始化; 条件表达式; 后置语句 {
		// do something
	}
------------------------------------------------------------------
for循环 初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。
for循环 一旦条件表达式的布尔值为 false，循环迭代就会终止。
------------------------------------------------------------------*/

package main

import "fmt"

func main() {

	// 输出从 0 到 9 数字
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v \t", i)
	}

	// 求和
	sum := 0
	for i := 10; i > 0; i-- {
		sum += i
	}
	fmt.Printf("\n\nsum: %v", sum)
}
