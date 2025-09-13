package catalog

import (
	"fmt"
	"net"

	"github.com/suryanshvermaa/microservices/golang/services/catalog/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedCatalogServiceServer
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterCatalogServiceServer(srv, &grpcServer{service: s})
	reflection.Register(srv)
	return srv.Serve(lis)
}
