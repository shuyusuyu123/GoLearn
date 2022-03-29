package main

import "fmt"

func main() {
	str :="hello,世界"
	n :=len(str)
	for i :=0;i<n;i++{
		ch :=str[i]
		fmt.Println("%d,%v",i,ch)
		//fmt.Printf(str)
	}
}