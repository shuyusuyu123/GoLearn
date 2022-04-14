package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	originStr := "aabccddddefffgha"
	//originStr := "aabccdAddDDCDDDDdefff"
	fmt.Println("originStr:\n", originStr)
	fileWriter(originStr)
	readStr := fileRead()
	fmt.Println("readStr:\n", readStr)
	compressStr := compressString(readStr)
	fileWriter(compressStr)
	fmt.Println("compressStr:\n", compressStr)
}

func fileWriter(str string) {
	file, err := os.OpenFile("samplefile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open fail:err=", err)
		return
	}
	defer file.Close()
	_, err = io.WriteString(file, str)
	io.WriteString(file, "\n")
	if err != nil {
		fmt.Println("write fail:err=", err)
		return
	}
}
func fileRead() string {
	data, err := os.ReadFile("samplefile.txt")
	if err != nil {
		fmt.Println("open fail:err=", err)
	}
	str := strings.Replace(string(data), "\n", "", -1)
	return string(str)
}
func compressString(str string) string {
	length := len(str)
	newStr := ""
	count := 0
	for i := 0; i < length; i++ {
		c := str[i]
		if i > 0 {
			cc := str[i-1]
			if c != cc {
				if count == 1 {
					newStr = newStr + string(cc)
					count = 1
				} else {
					newStr = newStr + string(cc) + strconv.Itoa(count)
					count = 1
				}

			} else {
				count++
			}
		} else {
			count++
		}
		if i == length-1 {
			if count == 1 {
				newStr = newStr + string(c)
			} else {
				newStr = newStr + string(c) + strconv.Itoa(count)
			}
		}
	}
	if length > len(newStr) {
		return newStr
	}
	return str
}
