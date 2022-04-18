package main

import (
	"fmt"
)

func main() {
	ss := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		ss = append(ss, i)
		//fmt.Printf("%d,", ss)
	}
	sum := 0
	for _, v := range ss {
		sum += v
	}
	fmt.Printf("%d", sum)
}
