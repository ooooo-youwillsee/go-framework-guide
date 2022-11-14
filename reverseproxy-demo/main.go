package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	reverseProxy := &httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL.Scheme = "https"
			request.URL.Host = "www.baidu.com"
			request.Host = "www.baidu.com"
		},
	}

	server := http.Server{
		Addr:    ":2222",
		Handler: reverseProxy,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
