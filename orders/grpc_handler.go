package main

import (
	"context"
	"log"

	pb "github.com/BharathMenon/commons/api"
	"google.golang.org/grpc"
)
type grpcHandler struct{
	pb.UnimplementedOrderServiceServer
}
func NewgrpcHandler(gs *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(gs,handler)
	
}
func (h* grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New Order Received!")
	o:=&pb.Order{
	ID:"42",
	}
	return  o,nil
}