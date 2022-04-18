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
	sumAll := []chan int{make(chan int), make(chan int), make(chan int), make(chan int)}
	num1 := ss[:len(ss)/4]
	num2 := ss[len(ss)/4 : len(ss)/2]
	num3 := ss[len(ss)/2 : len(ss)/4*3]
	num4 := ss[len(ss)/4*3:]
	go add(sumAll[0], num1...)
	go add(sumAll[1], num2...)
	go add(sumAll[2], num3...)
	go add(sumAll[3], num4...)
	sum := 0
	for _, c := range sumAll {
		sum += <-c
	}
	fmt.Printf("%d", sum)
}
func add(ch chan<- int, array ...int) {
	sum := 0
	for _, v := range array {
		sum += v
	}
	ch <- sum
	close(ch)
}
