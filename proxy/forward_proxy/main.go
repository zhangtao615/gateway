package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Pxy struct {}

// 正向代理
func (p *Pxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request %s %s %s\n", r.Method, r.Host, r.RemoteAddr)

	transport := http.DefaultTransport

	// 1. 浅拷贝对象，新增对象属性数据
	outReq := new(http.Request)
	*outReq = *r

	if clientIp, _, err := net.SplitHostPort(r.RemoteAddr); err != nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIp = strings.Join(prior, ", ") + ", " + clientIp
		}

		outReq.Header.Set("X-Forwarded-For", clientIp)
	}

	// 2. 请求下游
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	// 3. 把下游的内容返回给上游
	for key, value := range res.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}

	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	res.Body.Close()
}

func main() {
	fmt.Println("Serve on :8080")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}