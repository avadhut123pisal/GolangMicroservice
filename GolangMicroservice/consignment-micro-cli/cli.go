package main

import (
	pb "GolangMicroservice/consignment-service-micro/proto/consignment"
	// "google.golang.org/grpc"
	"log"
	"context"
	"os"
	"io/ioutil"
  "encoding/json"
  "github.com/micro/go-micro/cmd"
  microclient "github.com/micro/go-micro/client"
)

const (
	address = "localhost:8082"
	fileName = "consignment.json"
)

func parseFile(filename string) (*pb.Consignment, error){
  var consignment *pb.Consignment
  bytes, err:=ioutil.ReadFile(filename)
  if err!=nil {
	  return nil, err
  }
  json.Unmarshal(bytes, &consignment)
  return consignment, err
}

func main() {
  // set up a connection to server
  
  cmd.Init()
  
  
  client:=pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)
  file:=fileName
  if len(os.Args) > 1 {
	  file = os.Args[1]
  }
  consignment, err:=parseFile(file)
  resp, err:=client.CreateConsignment(context.Background(), consignment)
  if err!=nil {
	  log.Fatalf("Could not greet: %v", err)
  }
  log.Printf("Created: %t", resp.Created)
  
  // get consignments
  resp, err=client.GetConsignments(context.Background(), &pb.GetRequest{})
  if err!=nil{
	  log.Fatalf("Could not list consignments: %v", err)
  }
  for _, v:=range resp.Consignments{
	  log.Println(v)
  }
}