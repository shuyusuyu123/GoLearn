package main

import "fmt"

func CountChannel(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go CountChannel(chs[i])
	}
	for _, ch := range chs {
		<-ch
	}
}
