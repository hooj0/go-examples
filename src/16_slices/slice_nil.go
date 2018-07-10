// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-10
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
slice nil
==================================================================
切片默认值 为零值 nil
------------------------------------------------------------------
切片默认值 在一个空数组进行切片后，切片的值就是nil
切片默认值 nil 切片的长度和容量为 0 且没有底层数组
------------------------------------------------------------------*/
package main

import "fmt"

func main() {

	var arrays []int

	slices := arrays[:]
	fmt.Println("len: ", len(slices), ", cap: ", cap(slices), ", slices: ", slices)
	// len:  0 , cap:  0 , slices:  []

	if slices == nil {
		fmt.Println("slice is nil!")
	}
}
