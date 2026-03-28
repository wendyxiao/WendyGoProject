package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg1 sync.WaitGroup

func main() {
	//：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
	//另一个协程从通道中接收这些整数并打印出来。
	/*wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int) //无缓存
	go func() {          //匿名
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	//另外一个携程读取channel的数字
	go func() {
		defer wg.Done()
		//循环读取
		for val := range ch {
			fmt.Println(val)
		}
	}()
	wg.Wait()
	fmt.Println("end")*/
	//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	//考察点 ：通道的缓冲机制
	/*ch := make(chan int, 100)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println(value)
		}
	}()
	wg.Wait()*/
	//并发安全编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
	//每个协程对计数器进行1000次递增操作，最后输出计数器的值
	/*sc := safeCounter{
		mu:    sync.Mutex{},
		count: 0,
	}
	s := &sc
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				s.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf(strconv.Itoa(sc.count))*/
	//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
	//每个协程对计数器进行1000次递增操作， 最后输出计数器的值
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go incrementCounter()
	}
	wg1.Wait()
	fmt.Println("Counter:", counter)
}

type safeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *safeCounter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func incrementCounter() {
	defer wg1.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&counter, 1)
	}
}
