package client

import (
	"fmt"

	"github.com/olafszymanski/int-sdk/integration/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(host, port string) (pb.IntegrationClient, error) {
	// TODO: Secure the connection
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	cl := pb.NewIntegrationClient(conn)
	return cl, nil
}
