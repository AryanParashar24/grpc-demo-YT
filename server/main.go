package main 
import (
	"log"
	"net"
	pb "github.com/AryanParashar24/grpc-demo-YT/proto" // import the proto package where we defined the rpc rules	
	"google.golang.org/grpc"
)

const (
	port=":8080"	
)

type helloserver struct {
	pb.GreetServiceServer // implement the server interface from the proto file where we described about the rpc rules in the GreetService func
}

func main(){
	lis, err:= net.Listen("tcp", port) // to make the server listen and then follow the rpcs to make the connection
	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer() // create a new grpc server and since we doesnt have a log thatswhy we'll call by importing grpc
	pb.RegisterYourServiceServer(grpcServer, &helloserver{}) // register the service with the server
	log.Printf("server started at %v", lis.Addr()) // log the server address
	if err:= grpcServer.Serve(lis); err != nil { // start the server
		log.Fatalf("failed to serve: %v", err)
	}
}
