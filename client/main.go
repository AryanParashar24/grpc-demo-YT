package main 

import(
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure" // import the insecure package to use the insecure credential & this is used to create a grpc server without TLS, for development purposes
	pb "github.com/AryanParashar24/grpc-demo-YT/proto" // import the proto package where we defined the rpc rules	
)

const(
	port = ":8080" // port which we decided earleir in the server file on which the client will connect to the server

)

func main(){	// it just means that our client shoudl be able to create the ability to call our server which will  be done with the help of grpc.Dial() method
	conn, err:= grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredential))    // conn is the connection and its just using the localhost is 8080 without using any secure credentials thatswhy we kept it insecure.NewCredentials
	if err != nil {
		log.Fatalf("failed to connect to the server: %v", err) // if there is an error then log it
	}
	defer conn.Close() // close the connection when the function exits and server kept running with the client after it closed
	// thatswhy its been kept at the end of that the connection is closed after the execution is done at the end 

		client := pb.NewGreetServiceClient(conn) // create a new client using the connection and the service defined in the proto file

		// name:= &pb.NamesList{
		// 	Names: []string{"Aryan", "Parashar"}, // create a new NamesList with the names we want to greet
		// }

		callSayHello(client)	// we'll have to describe this function in orther to call this client
		

	}