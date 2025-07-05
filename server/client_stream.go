package main

import(
	"io"
	pb "github.com/AryanParashar24/grpc-demo-YT/proto" // import the proto package where we defined the rpc rules
)

func (s *HelloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error{
	var messages []string
	for{
		req, err:= stream.Recv() // receive the request from the client
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessagesList{Messages: messages}) // send the messages back to the client and close the stream when the end of file is reached
		}
		if err != nil{
			return err // return the error if there is an error
		}
		log.Printf("Got request with name: %v", req.Name)
		messages = append(messages, "Hello "+req.Name) // append the name to the messages slice
	}
}