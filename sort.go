package main

import (
	"fmt"
	"os"
	"strconv"
)

func sort(arr []int) {
	n := len(arr);
	for i:=0; i < n; i++ {
		for j:=1; j + i < n; j++ {
			if arr[j] < arr[j-1] {
				tmp := arr[j];
				arr[j] = arr[j-1];
				arr[j-1] = tmp;
			}
		}
	}
}

func biSearch(arr []int, v int) int {
	l := 0;
	h := len(arr) - 1;
	for l < h {
		m := (l + h + 1) / 2
		if arr[m] <= v {
			l = m
		} else {
			h = m - 1
		}
	}
	return l
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	v,err := strconv.Atoi(os.Args[1])
	if err != nil {
		print("Bad param~\n");
		return
	}
	arr := []int{18,1,45,22,333,25,72,37,168,19,52,92,29,12,89,62,97,28,54}
	sort(arr);
	fmt.Println(arr);
	n := biSearch(arr, v)
	fmt.Println(n);
}
