package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fullstorydev/grpchan"
	"github.com/fullstorydev/grpchan/httpgrpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"

	pb "httproxy/hello"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedHelloServer
}

// Say say
func (s *server) Say(ctx context.Context, in *pb.SayReq) (*pb.SayRsp, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.SayRsp{Reply: "Hello " + in.GetName()}, nil
}

var (
	port = flag.Int("port", 50057, "The server port")
)

func main() {

	svr := &server{}
	reg := grpchan.HandlerMap{}
	pb.RegisterHandlerHello(reg, svr)

	var mux http.ServeMux
	httpgrpc.HandleServices(mux.HandleFunc, "/", reg, nil, nil)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	httpServer := http.Server{Handler: &mux}
	go httpServer.Serve(lis)
	defer httpServer.Close()

	lis2, err := net.Listen("tcp", fmt.Sprintf(":%d", *port+1))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis2); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
