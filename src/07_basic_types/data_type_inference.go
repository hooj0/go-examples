// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
data type inference
==================================================================
数据类型推导
------------------------------------------------------------------
在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），
变量的类型由右值推导得出。
------------------------------------------------------------------
类型推导 程序根据右边数据推导变量类型
类型推导 赋值变量会根据变量赋值对象的引用进行推导
类型推导 变量在数字类型情况下，右边值的小数位决定推导类型
------------------------------------------------------------------*/


package main

import "fmt"

func main() {

	var k int
	j := k // j 也是一个 int

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("j: %v, i: %v, f: %v, g: %v \n", j, i, f, g) // j: 0, i: 42, f: 3.142, g: (0.867+0.5i)
	fmt.Printf("j: %T, i: %T, f: %T, g: %T \n", j, i, f, g) // j: int, i: int, f: float64, g: complex128
}