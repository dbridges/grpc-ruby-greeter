# gRPC Demo with Go and Ruby

This repo contains a demo gRPC Go server and Ruby client, utilizing Protobuf.

Before running you'll need to install the various gRPC libraries for Go and Ruby:

```
brew install protobuf
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u -v golang.org/x/tools/cmd/present
gem install grpc
gem install grpc-tools
```

Run the server:
```
go run server.go
```

Run the client:
```
ruby client.rb
```
