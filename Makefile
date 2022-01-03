
.PHONY: build-http-demo
build-http-demo:
	export GOPATH="$(realpath .)/experimental-code/http-demo/out"; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/ooooo-youwillsee/go-framework-guide/experimental-code/http-demo
