package main

import "fmt"

func main() {
	/*x := 10
	TestPoint(&x) //修改指针的值
	fmt.Println(x, "/n")
	a := 2
	b := 3
	swapValues(&a, &b)
	fmt.Println(a, b)*/
	s := []int{1, 2, 3}
	ChengyuTwo(&s)
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

// 指针：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// // 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
func TestPoint(p *int) {
	*p += 10 //解指针*x变成值
}

// 交换两个整数的值，使用指针
func swapValues(a, b *int) {
	*a, *b = *b, *a
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func ChengyuTwo(s *[]int) {
	for i := 0; i < len(*s); i++ {
		(*s)[i] *= 2
	}
}
