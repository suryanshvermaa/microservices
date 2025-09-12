package account

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct{
	service Service
}

func ListenGRPC(s Service,port int)error{
	list,err:=net.Listen("tcp",fmt.Sprintf(":%d",port))
	if err!=nil{
		return err
	}
	serv:=grpc.NewServer()
	pb.(serv,)
	reflection.Register(serv)
	return serv.Serve(list)
}

func (s *grpcServer) PostAccount(ctx context.Context,r *pb.) (*pb.,error){
	a,err:=s.service.PostAccount(ctx,r.Name)
	if err!=nil{
		return nil,err
	}
	return &pb.{}
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.) (*pb., error){
	a,err:=s.service.GetAccount(ctx,r.Id)
	if err!=nil{
		return nil,err
	}
	return &pb.{}
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.) (*pb., error){
	a,err:=s.service.GetAccounts(ctx,r.skip,r.take)
	if err!=nil{
		return nil,err
	}
	return &pb.{}
}