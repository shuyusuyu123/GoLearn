//package main
//
///*
//#include <stdlib.h>
//*/
//import "C"
//import "fmt"
//
//func Random() int {
//	return int(C.random())
//}
//
//func Seed(i int) {
//	C.srandom(C.uint(i))
//}
//func main() {
//	Seed(100)
//	fmt.Println("Random:", Random())
//}
package main

/*
#include <stdio.h>
void callC() {
   printf("Calling C code!\n");
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println("A Go statement!")
	C.callC()
	fmt.Println("Another Go statement!")
}
