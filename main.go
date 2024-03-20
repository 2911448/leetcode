package main

import (
	"fmt"
	"sync"
)

func main() {
	withChannel(5)
}

func withChannel(num int) {
	ch1, ch2 := make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	fmt.Println("start print:")

	go func() {
		defer wg.Done()
		for i := 1; i <= num; i += 2 {
			<-ch1
			fmt.Println("one: ", i)
			ch2 <- struct{}{}
		}
		<-ch1
		// 如果这里没有这一步，会直接死锁
	}()

	go func() {
		defer wg.Done()
		// for i := 2; i < num; i += 2 {  // 如果在这里没有= ，会直接死锁
		for i := 2; i <= num; i += 2 {
			<-ch2
			fmt.Println("two: ", i)
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{} // 启动的信号
	wg.Wait()
	fmt.Println("stop the work!")

}
