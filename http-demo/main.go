package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/test", test)
	serveMux.HandleFunc("/testEnv", testEnv)
	serveMux.HandleFunc("/testHeader", testHeader)

	log.Info("start http-server at http://localhost:2001")
	server := http.Server{Addr: ":2001", Handler: serveMux}
	if err := server.ListenAndServe(); err != nil {
		log.Errorln("start http-server err", err.Error())
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func testEnv(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	environ := os.Environ()
	for _, envKey := range environ {
		log.Infof("env: %s", envKey)
	}

	if key := r.Form["envKey"]; key != nil {
		w.Write([]byte(fmt.Sprintf("key: %s, value: %s", key[0], os.Getenv(key[0]))))
	}
}

func testHeader(writer http.ResponseWriter, request *http.Request) {
	if authorization := request.Header["Authorization"]; authorization != nil {
		writer.Header().Set("Authorization", authorization[0])
	}
	writer.Write([]byte("ok"))
}
