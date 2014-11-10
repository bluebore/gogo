package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("data.dat")
	if err != nil {
		fmt.Println("Create file fail")
		return
	}
	f.WriteString("hello go file operations")
	f.Close()

	f, err = os.Open("data.dat")
	if err != nil {
		fmt.Println("Open file fail")
	}
	defer f.Close()
	m := bufio.NewReader(f)
	line, isPrefix, err1 := m.ReadLine() //读行
	if err1 != nil {
		if err1 != io.EOF {
			err = err1
		}
		return
	}

	if isPrefix {
		fmt.Println("a too long line,seems unexpected.")
		return
	}

	str := string(line) //转换字符数组为字串
	fmt.Println(str)
}
