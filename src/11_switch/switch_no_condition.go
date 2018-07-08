// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-08
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
switch no condition
==================================================================
switch没有条件，一种 if-else if-else的变体用法
------------------------------------------------------------------
switch语句 可以没有条件参数语句，直接在case语句中完成判断语句
------------------------------------------------------------------*/
package main

import (
	"fmt"
	"time"
)

func main() {

	x := time.Now().Hour()

	// 没有条件参数
	switch {
		case x == 1:
			fmt.Println("x = 1")
		case x <= 3:
			fmt.Println("x <= 3")
		case x > 3:
			fmt.Println("x > 3")
	}
}