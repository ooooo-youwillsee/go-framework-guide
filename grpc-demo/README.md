# install protoc, go-plugin for compiler proto

https://grpc.io/docs/languages/go/quickstart/

# generate *_pb.go

```shell
cd grpc-demo
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```