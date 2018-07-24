// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-24
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
channel range close
==================================================================
通道关闭 发送者可以主动关闭通道，表示满意数据发送或接收
通道关闭 接受者 可以在接收通道数据，通过赋值表达式的第二个返回值判断是否有通道关闭
通道关闭 接受者 在判断没有值可以接收，且通道关闭，那么通道执行完成
------------------------------------------------------------------
语法：
	v, ok := <-ch

ok 通道是否关闭，true | false
------------------------------------------------------------------
通道迭代 会不间断的从通道中获取数据，直到通道被关闭
------------------------------------------------------------------
语法：
	for data := range channel
------------------------------------------------------------------
注意：
	只有发送者才能关闭信道，而接收者不能
	向一个已经关闭的信道发送数据会引发程序恐慌（panic）
	通道与文件不同，通常情况下无需关闭它们
	只有在必须告诉接收者不再有值需要发送的时候才有必要关闭，例如终止一个 range 循环。
------------------------------------------------------------------*/
package main

import "fmt"

func fibonacci(n int, channel chan int)  {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		channel <- x	// 发送数据
		x, y = y , x + y
	}

	close(channel)	// 关闭通道，才能停止 range 循环
}

func main() {

	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	// fibonacci(cap(ch), ch) // 不使用go routine 结果也没有问题

	for num := range ch {
		fmt.Println("num: ", num)
	}
}
