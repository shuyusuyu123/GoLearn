package main

import (
	"fmt"
)

func main() {
	//learnString()
	//learnSlice()
	//learnFormat()
	learnPrint()
}

//变量
func learnVar() {
	var a int = 1
	var b int = 2 * 3
	var c int = b
	var d int
	e, f := 10, 20.1
	g := 20.1
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
	fmt.Println("f=", f)
	fmt.Println("g=", g)
}

//常量
const (
	MathB0 = 1 << iota
	MathB1 = 1 << iota
	MathB2 = 1 << iota
)
const (
	MathA0 = iota
	MathA1
	MathA2
)

//BOOL类型
func learnBool() {
	a := false
	fmt.Println(a)
	a = true
	fmt.Println(a)
}

//字符串以及字符串字串
func learnString() {
	//=========string to byte=============
	//string in const
	str := "hello,this is helen."
	//conver str to byte slice
	//warning: this will copy new block memory
	byteStr := []byte(str)

	fmt.Println(str)
	fmt.Println(byteStr)
	//str[1]="c"//build fail
	byteStr[1] = 'c'
	fmt.Println(byteStr)
	//================string to sub string==========
	suba := str[6:10]
	subb := str[6:]
	subc := str[:20]
	fmt.Println(suba)
	fmt.Println(subb)
	fmt.Println(subc)
}
func learnStringPlus() {
	a := "hello,你好"
	b := []byte(a)
	c := []rune(a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	for i, v := range a {
		fmt.Println(i, v)
	}

}

//指针类型
func learnPointer() {
	var a = 11
	p := &a
	fmt.Println("a=", a, "p=", p, "*p=", *p)
}

//结构体类型
func learStruct() {
	type User struct {
		name string
		age  int
	}
	Helen := &User{
		name: "yshu",
		age:  18,
	}
	fmt.Println("Helen.name=", Helen.name)
	fmt.Println("Helen.age=", Helen.age)
}

//动态数组
func learnSlice() {
	var a []int
	fmt.Println(a, ",len(a)=", len(a), ",cap(a)=", cap(a))
	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(a, ",len(a)=", len(a), ",cap(a)=", cap(a))
	b := make([]int, 10, 100)
	fmt.Println(b, ",len(b)=", len(b), ",cap(b)=", cap(b))
	copy(b, a)
	fmt.Println(b, ",len(b)=", len(b), ",cap(b)=", cap(b))
	b = append(a, a...)
	fmt.Println(b, ",len(b)=", len(b), ",cap(b)=", cap(b))
	//for range
	for i, v := range b {
		fmt.Println(i, v)
	}
}

//Map
func learnMap() {
	mapA := make(map[int]string)
	fmt.Println("len(mapA)=", len(mapA))
	mapA[10] = "0001"
	mapA[20] = "0002"
	mapA[90] = "0003"
	mapA[111] = "0004"
	fmt.Println(mapA)
	fmt.Println("len(mapA)=", len(mapA))
	for k, v := range mapA {
		fmt.Println(k, v)
	}
	//Query
	v := mapA[90]
	fmt.Println("v=", v)
	//Delete
	delete(mapA, 90)
	//OK,Query
	value, ok := mapA[90]
	fmt.Println("value=", value, ",ok=", ok)

	value, ok = mapA[20]
	fmt.Println("value=", value, ",ok=", ok)

}

//Format
func learnFormat() {
	fmt.Println("%v\n", "hello")
	fmt.Println("%q\n", "hello")
}

//learn all print format here.
//https://www.jianshu.com/p/8be8d36e779c
func learnPrint() {
	a := 0
	b := 0.9
	c := "abc"
	fmt.Println("a=%v b=%v c=%v\n", a, b, c)
	type Person struct {
		age  int
		name string
	}
	d := &Person{age: 9, name: "Helen"}
	fmt.Println("%v\n", d)
	fmt.Println("%+v\n", d)
}
