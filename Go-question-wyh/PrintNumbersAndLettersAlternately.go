package main

import "fmt"

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/2/29 2:48 下午
 */

//使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个goroutine打印字母
//最终效果如下:
//12AB34CD56EF78GH910IJ

func main() {
	numbers := make(chan bool, 1)
	letter := make(chan bool, 1)
	done := make(chan int)

	go func() {
		for i := 1; i < 11; i += 2 {
			<-letter
			fmt.Print(i)
			fmt.Print(i + 1)
			numbers <- true
		}
	}()

	go func() {
		sequence := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
		for i := 0; i < 10; i += 2 {
			<-numbers
			fmt.Print(sequence[i])
			fmt.Print(sequence[i+1])
			letter <- true
		}
		done <- 1
	}()

	letter <- true
	<-done
}
