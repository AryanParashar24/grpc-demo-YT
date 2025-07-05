package main

import (
	
	pb "github.com/AryanParashar24/grpc-demo-YT/proto" // import the proto package where we defined the rpc rules
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, name *pb.NamesList){
	log.Printf("Bidirectional Streaming started")
	stream, err:= client.SayHelloBiDirectionalStream(context.Background()) // call the SayHelloBiDirectionalStream method from the GreetServiceClient interface with the context and the request
	if err != nil{
		log.Fatalf("could not send names: %v", err) // if there is an error then log it
	}

	waitc:= make(chan, struct{})	// here we created a channel as we'll have to manage the stream of message being sent and been received in the response from the server

	go func(){
		for {
			res, err:= stream.Recv() // receive the response from the server
			if err == io.EOF{	// check if the file havent ended 
				break
			}
			if err != nil{
				log.Fatalf("error while streaming: %v", err)
			}
			log.Println(res.Message) // print the message received from the server
		}
		close(waitc) // close the channel when the stream is done
	}()

	for _, name := range name.Names { // iterate over the names in the request and then creating the objects out of it. 
		re:= &pb.HelloRequest{	// here HelloRequest is a struct whihc requires name
			Name: name, // set the name in the request
		}	// now we will need to receive a stream of response from the server to the client and a stream of requests to the server from the clients
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending request: %v", err) 
		}
		time.Sleep(2 * time.Second) // sleep for 2 seconds before sending the next message
	}
	stream.CloseSend() // close the stream when done sending all the messages
	<-waitc // wait for the channel to be closed before exiting the function
	log.Printf("Bidirectional Streaming finished")
}