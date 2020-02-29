package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/2/29 4:35 下午
 */

func main() {
	output := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			output <- rand.Int()
		}
		close(output)
	}()
	go func() {
		defer wg.Done()
		for i := range output {
			fmt.Println(i)
		}
	}()
}
