package main

import (
	
	pb "github.com/BharathMenon/commons/api"	
	"context"
)

type OrdersService interface{
	CreateOrder(context.Context) error
	ValidateOrder(ctx context.Context,co *pb.CreateOrderRequest) error
 }

 type OrderStore interface{
	Create(context.Context) error
 }