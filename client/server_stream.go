package main 
import(
	"context"
	pb "github.com/AryanParashar24/grpc-demo-YT/proto" // import the proto package where we defined the rpc rules
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("Streaming started")
	stream, errr:= client.SayHelloServerStreaming(contecxt.Background(), names) // call the SayHelloServerStreaming method from the GreetServiceClient interface with the context and the request
	if errr != nil{ 
		log.fatalf("could not send names: %v", errr) // if there is an error then log it
	}
	// now we know that the client will send a req to which server will eb streaming the messages and the client will be receiving it and processing it & then print it
	for {
		message, err := stream.Recv()
		if err == io.EOF{	// check if the file havent ended 
			break
		}
		if err != nil{
			log.fatal("error whiile streaming: %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming ended") // log the end of the streaming
}
