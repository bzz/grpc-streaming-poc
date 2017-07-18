Server-side gRPC streaming in Glong
------

```
#install `protoc`
go get -u github.com/gogo/protobuf/...
protoc --go_out=plugins=grpc:./streaming streaming.proto

go run server.go
go run clinet
```


TODO:
 - build an exmaple .go file \w annotations for Proteus and an interface proposal
   .Next() ? chan(*)?
 - create issue \w feature request in proteus repo