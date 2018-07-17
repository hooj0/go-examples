// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
struct
==================================================================
结构体 一种多个不同（相同）数据集合对象，相当于一个实体对象模型
------------------------------------------------------------------
语法：
	type 名称 struct {
		变量属性..1
		变量属性..2
		变量属性..N
	}
------------------------------------------------------------------
结构体 可以声明一个复杂的对象类型，来存储各种复杂的数据
结构体 可以作为各个数据层之间数据处理的对象模型，作为传递数据用极好
结构体 便于维护管理，可读性好，移植性好
结构体 即便没有初始化，也是有零值，其中属性的零值对应其类型的默认值
------------------------------------------------------------------*/
package main

import "fmt"

type Vertex struct {
	x int
	y int
}


func main() {

	// 初始化 Vertex 结构体，相当于创建一个Vertex实例
	var vertex = Vertex{ 2, 5 }
	// 直接读取结构体对象值
	fmt.Println(vertex)

	// 访问 Vertex 结构体中的属性数据
	fmt.Println("x: ", vertex.x, "y: ", vertex.y)

	// 再次实例化赋值
	vertex = Vertex{ 11, 30 }
	fmt.Println("x: ", vertex.x, "y: ", vertex.y)

	// 单个属性赋值
	vertex.y = 666
	fmt.Println("x: ", vertex.x, "y: ", vertex.y)
}
