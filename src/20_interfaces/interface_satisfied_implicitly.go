// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-16
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
interface satisfied implicitly
==================================================================
接口隐式实现 类型可以实现一个接口的所有方法来完成接口实现
------------------------------------------------------------------
接口隐式实现 没有显示的实现关键字 implements
接口隐式实现 具有解耦功能，可以在任何包任何地方进行实现，从而实现也无需增加接口名称
------------------------------------------------------------------

------------------------------------------------------------------*/
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事
func (t T) M() {
	fmt.Println(t.S)
}

func main() {

	var i I = T{ "interface implements" }
	i.M()
}
