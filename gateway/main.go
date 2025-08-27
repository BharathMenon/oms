package main

import (
	"log"
	"net/http"

	"github.com/BharathMenon/commons"
	pb "github.com/BharathMenon/commons/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var(
	httpAddr = commons.EnvString("HTTP_ADDR",":3000")
	orderServiceaddr = "localhost:3000"
)

func main(){
	conn, err := grpc.NewClient(orderServiceaddr,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("Failed to dial server: %v",err)
	}
	defer conn.Close() 
	log.Println("Dialing orders service at ",orderServiceaddr)
	c := pb.NewOrderServiceClient(conn)
	mux := http.NewServeMux()
	
	handler:= NewHandler(c)
	handler.registerRoutes(mux)
	log.Printf("Starting HTTP Server at %s",httpAddr)
	if err:= http.ListenAndServe(httpAddr,mux); err!=nil{
		log.Fatal("Failed to start http server")
	}

}
