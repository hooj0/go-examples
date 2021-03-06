// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-17
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
interface nil
==================================================================
nil接口值 没有具体实现的接口，即没有具体值也没有具体实现类型
------------------------------------------------------------------
nil接口值 因为没有具体实现类型，接口不知道调用那个实现，故会出错
------------------------------------------------------------------
------------------------------------------------------------------*/
package main

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {

	var i I
	// 接口元组数据也无法知道具体的实现类型
	describe(i)		// (<nil>, <nil>)
	i.M()			// panic: runtime error: invalid memory address or nil pointer dereference
}
