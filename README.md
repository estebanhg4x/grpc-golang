# grpc-golang



# generate package 
protoc --go_out=. proto/hellopb.proto
# plugins  grpc

protoc --go_out=. --go-grpc_out=. proto/hello.proto


# AMBOS Plugins

protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/hello.proto


// mustEmbedUnimplementedHelloServiceServer implements proto.HelloServiceServer.
func (*server) mustEmbedUnimplementedHelloServiceServer() {
	panic("unimplemented")
}