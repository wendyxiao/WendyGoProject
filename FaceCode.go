package main

import "fmt"

func main() {
	/*	r := rectange{

			4.01,
			1.01,
		}
		fmt.Println(r.area())
		fmt.Println(r.perimerter())
		c := circle{
			r: 8,
		}
		fmt.Println(c.area())
		fmt.Println(c.perimerter())*/
	e := Employee{
		EmployeeId: 1,
		p: Person{
			Name: "wendy",
			Age:  38,
		},
	}
	e.printInfo()
}

// 定义接口
type shape interface {
	//计算面积
	area() float64
	perimerter() float64
}

// 实现该接口
type rectange struct {
	with   float64
	height float64
}

func (r rectange) area() float64 {
	return r.with * r.height
}
func (r rectange) perimerter() float64 {
	return (r.with + r.height) * 2
}

// 实现接口shape
type circle struct {
	r float64
}

func (c circle) area() float64 {
	return c.r * c.r * 3.14
}
func (c circle) perimerter() float64 {
	return c.r * 3.14 * 2
}

//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//考察点 ：组合的使用、方法接收者

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	EmployeeId int
	p          Person
}

// 打印员工信息
func (e Employee) printInfo() {
	fmt.Println("employeeId", e.EmployeeId)
	fmt.Println("name:", e.p.Name)
	fmt.Println("age:", e.p.Age)
}
