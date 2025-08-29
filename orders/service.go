package main

import (
	"context"
	"log"
	

	"github.com/BharathMenon/commons"
	pb "github.com/BharathMenon/commons/api"
)


type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context)error{
	return nil
}

func (s *service) ValidateOrder(ctx context.Context,co *pb.CreateOrderRequest) error{
  if len(co.Items) ==0{
	return commons.ErrNoItems
  }


  mergedItems := mergeItemQuantities((co.Items)) 
  log.Print(mergedItems)
  //Now validate with stock service
  return nil
	return nil
}

func mergeItemQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity{
	merged:= make([]*pb.ItemsWithQuantity,0)
	for _,item := range items{
		found:=false
		for _,fitem := range merged{
			if fitem.ID==item.ID{
				fitem.Quantity+=item.Quantity
				found = true
				break
			}
		}
		if !found{
			merged = append(merged, item)
		}
	}
	return merged
}