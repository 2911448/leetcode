package main

import (
	"fmt"
	"sync"
)

func main() {
	withChannel()
}

func withChannel() {
	ch1, ch2 := make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	fmt.Println("start print:")

	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-ch1
			fmt.Println(i)
			ch2 <- struct{}{}
		}
		<-ch1
		// 如果这里没有这一步，会直接死锁
	}()

	go func() {
		defer wg.Done()
		for i := 65; i <= 90; i++ {
			<-ch2
			fmt.Println(string(byte(i)))
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{} // 启动的信号
	wg.Wait()
	fmt.Println("stop the work!")

}
