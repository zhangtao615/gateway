package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.DefaultClient.Get("http://localhost:8080/ping")
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	bds, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bds))
}