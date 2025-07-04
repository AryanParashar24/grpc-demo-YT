package main
import(
	"context"
	"time"
	pb "github.com/AryanParashar24/grpc-demo-YT/proto" // import the proto package where we defined the rpc rules

)
//as we have defined int he proto file that the SayHello method will not take any parameters and will return a response with the greeting message
// now as described earlier in the unary context the speciofications r mentioned in the unary file in server and one at client side amongst which
// this unary file will be a calling back function where we will be sending a noParam request with a context from to the server in order to get a response which will be shown in the last line after executing the log.Printf 

func callSayHello(client pb.GreetServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // create a context with a timeout of 1 second
	defer cancel() // cancel the context when the function exits


	res, err:= client.SayHello(ctx, &pb.NoParam{}) // call the SayHello method from the GreetServiceClient interface with the context and the request
	if err != nil{
		log.Fatalf("could not greet: %v", err) // if there is an error then log it
	}
	log.Printf("%s", res.Message)	// here theses messages a r from the proto file where we described about the rpc rules adn the messages
}