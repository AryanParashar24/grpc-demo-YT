package main
 
import (
	"context"
	pb "github.com/AryanParashar24/grpc-demo-YT/proto" // import the proto package where we defined the rpc rules
)

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.GreetResponse, error) {
	// implement the helloserver method from the GreetServiceServer interface and the function is called SayHello with context.Context as the context and the pb.NoParam as the request to the SayHello function from GreetResponse interface
	// this method will be called by the client when it wants to greet the server & we will return a response with the greeting message

	// Now the unary specififcations are been specified in the unary files in the server and the client side amongst whihc the server side unary file is supposed to response a Hello message
	// and the client side unary file is supposed to call the SayHello method from the GreetServiceClient interface with the context and the request// while in the client unary that will be a calling back function where we will be sending a request from 
	return &pb.HelloResponse{
		Message: "Hello", 	// response that has to be passed from the server to the client in order to form up the unary connection between the server and the client 
	}, nil
}