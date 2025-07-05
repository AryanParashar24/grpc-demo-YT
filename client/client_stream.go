package main

import(
	"time"
	"context"
	pb "github.com/AryanParashar24/grpc-demo-YT/proto" // import the proto package where we defined the rpc rules
)

func callSayHelloClientStream(client pb.GreetServiceClient, name *pb.NamesList){
	log.Printf("Client streaming started")
	stream, err:= client.SayHelloClientStreaming(context.Background())

	if err != nil{
		log.Fatalf("could not send names: %v", err) // if there is an error then log it
	}	
	for _, n := range name.Names {
		req := &pb.Hellorequest{	// as we know HelloRequest is a struct 
			Name: name, 
		}
		if err:= stream.Send(req); err != nil { // checked if there is an error while sending the request
			log.fatal("Enter while sending: %v", err)
		}
	log.Printf("Send the request with name: %v", name)
	time.Sleep(2*time.Second)
	}

	res, err:= stream.CloseAndRecv() // close the stream and receive the response from the server
	log.Printf("Client streaming finished")
	if err != nil {
		log.fatalf("error while receiving %v", err)
	}
	log.Printf("%v", resMessages)
}
