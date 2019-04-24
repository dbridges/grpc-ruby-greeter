.PHONY: proto

proto:
	protoc -I . greeter.proto --go_out=plugins=grpc:greeter
	grpc_tools_ruby_protoc -I . --ruby_out=. --grpc_out=. greeter.proto
