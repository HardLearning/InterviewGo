package main

import (
	"fmt"
	_ "runtime"
)

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/2/29 10:52 下午
 */

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	for value := range ch {
		fmt.Print(value)
	}
}
