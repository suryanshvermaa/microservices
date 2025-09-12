package account

import "context"

type grpcServer struct{
	service Service
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