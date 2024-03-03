package server

import (
	"fmt"
	"net"

	"github.com/olafszymanski/int-sdk/integration/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedIntegrationServer
}

func Start(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterIntegrationServer(s, &server{})
	return s.Serve(lis)
}
