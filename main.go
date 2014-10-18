package main

import "fmt"

var name = "bluebore"

func func1() {
	print("call func1\n");
	print(name)
	fmt.Println("\n")
}

const PI = 3.14

type node struct{
	left *node;
	right *node;
	value string;
}

type golang interface{}

func main() {
	print("Hello golang\n")
	func1()
}
