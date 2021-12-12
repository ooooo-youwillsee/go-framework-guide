package main

import (
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"os"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/test", test)
	serveMux.HandleFunc("/testEnv", testEnv)
	serveMux.HandleFunc("/testHeader", testHeader)

	server := http.Server{Addr: ":2001", Handler: serveMux}
	server.ListenAndServe()
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func testEnv(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	environ := os.Environ()
	for _, envKey := range environ {
		glog.Infof("env: %s", envKey)
	}

	if key := r.Form["envKey"]; key != nil {
		w.Write([]byte(fmt.Sprintf("key: %s, value: %s", key[0], os.Getenv(key[0]))))
	}
}

func testHeader(writer http.ResponseWriter, request *http.Request) {
	if aaa := request.Header["Authorization"]; aaa != nil {
		writer.Header().Set("Authorization", aaa[0])
	}
	writer.Write([]byte("ok"))
}
