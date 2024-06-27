package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "serverGRPC/server"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

type Data struct {
	Texto string `json:"texto"`
	Pais  string `json:"pais"`
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {

	tweet := map[string]string{
		"texto": in.GetTexto(),
		"pais":  in.GetPais(),
	}

	fmt.Println("El cliente recibio el album: ", tweet)

	return &pb.ReplyInfo{Info: "Hola client, recibi el album"}, nil

}
func main() {

	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		panic(err)
	}

}
