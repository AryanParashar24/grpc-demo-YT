package main
import(
	"log"
	"time"
	pb "github.com/AryanParashar24/grpc-demo-YT/proto"
)

func (s *HelloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	// implement the SayHelloServerStream method from the GreetServiceServer interface
	// this method will be called by the client when it wants to greet the server with a list of names
	log.printf("got requests with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResopnse{
			Message: "Hello" + name, 
		}
		if err:= stream.Send(res); err != nil {
			return err // return the error if there is an error
		}
		time.Sleep(2*time.Second)
	}
	return nil // return nil if there is no error
}