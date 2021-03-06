// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-09
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
array
==================================================================
数组 用来有序的存取多个相同类型数据的集合
------------------------------------------------------------------
数组 存储一组相同类型的数据
数据 存储可以按照放入的顺序固定
数据 可以快速按照索引读取数据
数组 可以快速按照特定位置插入数据
数组 有着固定大小空间的集合，不可改变的数据集合
------------------------------------------------------------------
数组是值，将一个数组赋予另一个数组会复制其所有元素。
若将某个数组传入某个函数，它将接收到该数组的一份副本而非指针。
数组的大小是其类型的一部分。类型 [10]int 和 [20]int 是不同的。
------------------------------------------------------------------
数组为值的属性很有用，但代价高昂；若你想要C那样的行为和效率，你可以传递一个指向该数组的指针

func Sum(a *[3]float64) (sum float64) {
	// do something...
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // 注意显式的取址操作
------------------------------------------------------------------*/
package main

import "fmt"

func main() {

	// 声明一个长度为2的字符串类型的数组
	var strings [2]string
	strings[0] = "Go"
	strings[1] = "Lang"
	fmt.Println("strings[0]: ", strings[0])
	fmt.Println("strings: ", strings)

	// 声明一个数组，给出默认值和固定长度
	var numbers  = [5]int { 1, 3, 2, 4, 5 }
	fmt.Println("numbers: ", numbers)

	// 声明一个字符串数组，没有固定长度
	names := []string { "jack", "tom" }
	fmt.Println("names: ", names)

	// 由于没有设置固定大小，在构建数组期间就默认设置元素个数就是数组大小，所以后期不能超越数组大小设置元素
	// names[3] = "jason"
	// fmt.Println("names: ", names)
}
