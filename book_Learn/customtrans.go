package main

import (
	"fmt"
	"net/http"
)

type OurCustonTransport struct {
	Transport http.RoundTripper
}

func (t *OurCustonTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *OurCustonTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	//处理一些事情
	//发起HTTP请求
	//添加一些域到req.Header中
	return t.transport().RoundTrip(req)
}
func (t *OurCustonTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}
func main() {
	t := &OurCustonTransport{
		//...
	}
	c := t.Client()
	resp, err := c.Get("http://www.baidu.com")
	fmt.Println("resp", resp.Body, "::err", err)
	//...
}
