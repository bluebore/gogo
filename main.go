package main

import "fmt"

func maopao() {
	a := [20]int{18,1,45,22,333,25,72,37,168,19,52,92,29,12,89,62,97,28,54}
	fmt.Println(a);
	n := len(a);
	for i := 0; i<n; i++ {
		for j := 1; j+i<n; j++ {
			if a[j] < a[j-1] {
				x := a[j];
				a[j] = a[j-1]
				a[j-1] = x;
			}
		}
	}
	fmt.Println(a);
}
var name string = "bluebore"

func func1() {
	print("call func1\n");
	print(name)
	fmt.Println("\n")
}

const (
	PI = 3.14
)

const (
	Monday = iota
	Tuesday
	Wednesday
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	EB
	ZB
)

type node struct{
	left *node;
	right *node;
	value string;
}

type golang interface{}

func main() {
	print("Hello golang\n")
	func1()
	var a[3]int
	fmt.Println(a)

	// 循环枚举
	for i := B; i <= 2*B; i++ {
		fmt.Println(i)
	}

	// 指针
	x := 88
	var p *int = &x
	fmt.Println(*p)
	
	if a := 8; a < 10 {
		fmt.Println("Run if");
	}
	
	var ar [20]int = [20]int{18:18,19}
	par := new([20]int);
	fmt.Println(ar)
	fmt.Println(par)
	par = &ar
	fmt.Println(par)

	vv := [2][3]int{
			{2,3,4},
			{1,1,1}}
	fmt.Println(vv);
	maopao()

	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1);
	s1 = append(s1, 1,2,3,4);
	fmt.Printf("%p\n", s1);
	
	var m map[int]string = make(map[int]string);
	m[3] = "OK3"
	m[1] = "OK1"
	m[2] = "OK2"
	fmt.Println(m)
	for _,v:=range m{
		fmt.Println(v)
	}
	delete(m, 1);
	fmt.Println(m)
}
