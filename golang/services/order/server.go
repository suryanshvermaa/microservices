package order

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/suryanshvermaa/microservices/golang/services/account"
	"github.com/suryanshvermaa/microservices/golang/services/catalog"
	"github.com/suryanshvermaa/microservices/golang/services/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service       Service
	accountClient *account.Client
	catalogClient *catalog.Client
}

func ListenGRPC(s Service, accountURL, catalogURL string, port int) error {
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return err
	}
	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return err
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterOrderServiceServer(serv, &grpcServer{
		s,
		accountClient,
		catalogClient,
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostOrder(ctx context.Context, r *pb.PostOrderRequest) (*pb.PostOrderResponse, error)
{
	_,err:=s.accountClient.GetAccount(ctx,r.AccountId)
	if err != nil {
		log.Println("error in getting account:",err)
		return nil, err
	}
}