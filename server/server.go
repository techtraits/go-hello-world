package server

import (
	"fmt"
	"log"
	"net"

	pb "github.com/techtraits/go-hello-world/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"math/rand"
	"time"
)

var colours = [...]string{"CadetBlue", "Brown", "CornflowerBlue", "DarkGoldenRod", "DarkKhaki", "DarkOliveGreen", "DarkOrange"}

type greeterService struct {
	//TODO: MySQL This
	counters map[string]uint32
	colour   string
}

// Return the hello greeting
func (g *greeterService) SayHello(ctx context.Context, user *pb.User) (*pb.Greeting, error) {

	counter := g.counters[fmt.Sprintf("%s_%s", user.FirstName, user.LastName)]
	counter = counter + 1
	g.counters[fmt.Sprintf("%s_%s", user.FirstName, user.LastName)] = counter
	return &pb.Greeting{Text: formatMessage(user.FirstName, user.LastName, counter, g.colour)}, nil
}

func formatMessage(firstName string, lastName string, counter uint32, colour string) string {
	//<font color="red">This is some text!</font>
	message := fmt.Sprintf("Hello %s %s, I have seen you %d times",
		firstName, lastName, counter)
	return fmt.Sprintf("<font color=\"%s\">%s</font>", colour, message)
}
func RunServer(port int32) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Int31n(7)
	log.Printf("Color %d", index)
	pb.RegisterGreeterServer(grpcServer, &greeterService{make(map[string]uint32), colours[index]})
	log.Printf("Starting Greeter Server on port %d", port)
	grpcServer.Serve(lis)
}
