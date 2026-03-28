package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	//编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	//考察点 ： go 关键字的使用、协程的并发执行。
	/*	wg.Add(2)
		go printQiNumber()
		go printOuNumber()
		wg.Wait()*/
	//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	for i := 0; i < 100; i++ {
		go quarterJob(i)
	}
	time.Sleep(1 * time.Second)
}
func quarterJob(i int) {
	//print current time
	fmt.Println("quarter job start", time.Now())
	i++
	fmt.Printf("quarter job %d\n", i)
	fmt.Println("quarter job end", time.Now())
}
func printQiNumber() {
	defer wg.Done()
	for i := 1; i < 10; i += 2 {
		fmt.Println(i)
	}
}
func printOuNumber() {
	defer wg.Done()
	for i := 2; i < 10; i += 2 {
		fmt.Println(i)
	}
}
