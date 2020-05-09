# gRPC

## Plugins

```
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
```

## Create pb files

```bash
$ protoc --go_out=plugins=grpc:./pb ./protos/*.proto
```

## Create Docs

```bash
$ protoc --doc_out=html,index.html:./docs ./protos/*.proto
```
