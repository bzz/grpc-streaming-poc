
```
#install `protoc`
go get -u github.com/gogo/protobuf/...
protoc --go_out=plugins=grpc:./streaming streaming.proto

go run server.go
go run clinet
```


TODO: build exmaple .go file \w annotations for Proteus and interface proposal