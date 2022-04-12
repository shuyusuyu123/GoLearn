package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{12, 7, 11, 15}
	target := 23
	fmt.Println(find(arr, target))
}

func find(arr []int, target int) []int {
	sort.Ints(arr)
	myResult := make([]int, 0, 2)
	i := 0
	j := len(arr) - 1
	for {
		if arr[i]+arr[j] == target {
			myResult = append(myResult, arr[i], arr[j])
			break
		} else if arr[i]+arr[j] > target {
			j--
		} else {
			i++
		}
		if i >= j {
			break
		}
	}
	return myResult
}
