// --------------------------------------------------------------
// author: hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create date: 2018-06-28
// copyright by hoojo @ 2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
func
==================================================================
函数 一个功能单元，函数可以完成一段业务，让程序可以重用
------------------------------------------------------------------
函数 对程序单独的功能单元进行分解，让程序更有可读性、条例更清晰
函数 可以重复运行，一般在程序中，多次出现的代码可以作为一个函数被复用
函数 可以接收多个参数，返回多个参数
------------------------------------------------------------------*/

package main

func add(x int, y int) int {
	return x + y
}

func main() {

	println("add number: ", add(2, 3))
}