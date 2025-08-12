package main

import "context"

type store struct {
	//add mongodb instance here
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context){
	//
}