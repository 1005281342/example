package main

import (
	"context"
	"google.golang.org/grpc"
	pb "httproxy/hello"
	"log"
)

func main() {
	conn, err := grpc.Dial(":50058", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	var (
		cli = pb.NewHelloClient(conn)
		rsp *pb.SayRsp
	)
	if rsp, err = cli.Say(context.Background(), &pb.SayReq{Name: "me"}); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", rsp)

}
