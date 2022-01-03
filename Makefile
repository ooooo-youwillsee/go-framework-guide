
build-http-demo :
	cd experimental-code/http-demo
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/http-demo main.go
