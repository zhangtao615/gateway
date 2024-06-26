package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

var (
	proxy_addr = "http://127.0.0.1:2003"
	port       = "2002"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析代理地址，更改请求体的协议和主机
	proxy, err := url.Parse(proxy_addr)

	if err != nil {
		log.Println(err)
		return
	}
	
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host

	// 2. 请求下游
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(r)
	if err != nil {
		log.Println(err)
		return
	}

	// 3 把下游请求内容返回给上游
	for key, value := range resp.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(w)
}

// http 反向代理
func main() {
	http.HandleFunc("/", Handler)
	log.Println("start serving on port " + port)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Fatal(err)
	}
}