package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErrorHttp2(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErrorHttp2(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErrorHttp2(err)

	result, err := ioutil.ReadAll(conn)
	checkErrorHttp2(err)
	fmt.Println(string(result))
	os.Exit(0)
}
func checkErrorHttp2(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal erroe:%s", err.Error())
		os.Exit(1)
	}
}
