package main

import "context"

type store struct {
	//add mongodb instance here
	os *OrderStore
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context)error{
	//
	return nil
}