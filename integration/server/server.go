package server

import (
	"fmt"
	"net"

	"github.com/olafszymanski/int-sdk/integration/pb"
	"google.golang.org/grpc"
)

func Start(integration pb.IntegrationServer, port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterIntegrationServer(s, integration)
	return s.Serve(lis)
}
