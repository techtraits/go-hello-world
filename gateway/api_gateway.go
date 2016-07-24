package gateway

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	pb "github.com/techtraits/go-hello-world/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var shost string
var client pb.GreeterClient

func RunGateway(port int32, service_host string) error {

	e := echo.New()
	e.GET("/", getGreeting)
	shost = service_host
	log.Printf("Connecting to backend %s and running on %d", shost, port)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(shost, opts...)
	defer conn.Close()

	if err != nil {
		log.Printf("Error conecting to server: %s", err.Error())
		return err
	} else {
		log.Printf("Connected to Server")
		client = pb.NewGreeterClient(conn)
	}

	e.Run(standard.New(fmt.Sprintf(":%d", port)))
	log.Printf("Closing gateway service")
	return nil
}

func getGreeting(c echo.Context) error {

	first_name := c.QueryParam("first_name")
	last_name := c.QueryParam("last_name")

	if first_name == "" || last_name == "" {
		c.HTML(http.StatusBadRequest, "Missing required parameter first_name or last_name")
	} else {
		greeting, err := client.SayHello(context.Background(), &pb.User{FirstName: first_name, LastName: last_name})
		if err != nil {
			log.Printf("Error calling helloworld: %s", err.Error())
			return err
		} else {
			log.Printf("%s", greeting.Text)
			c.HTML(http.StatusOK, greeting.Text)
		}
	}
	return nil
}
