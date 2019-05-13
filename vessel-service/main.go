package main

import (
	"os"
	"fmt"
	"github.com/avadhut123pisal/GolangMicroservice/vessel-service/proto/vessel"
	// pb "GolangMicroservice/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
	"log"
)

const (
	DEFAULT_DB_HOST = "localhost:27017"
)

func main() {
   uri:=os.Getenv("DB_HOST")
   if uri == "" {
	 uri = DEFAULT_DB_HOST
   }
   
   // Mgo creates a 'master' session, we need to end that session
   // before the main function closes.
   session, err:=CreateSession(uri)
   defer session.Close()
   if err != nil {
     log.Panicf("Could not connect to datastore with host %s - %v", uri, err)
   }
   repo:=&VesselRepository{session}
   
   // create some dummy vessel data
   createDummyData(repo)
   
   srv:=micro.NewService(
	   micro.Name("go.micro.srv.vessel"),
	   micro.Version("latest"),
   )
   
   srv.Init()
   
   pb.RegisterVesselServiceHandler(srv.Server(), &handler{repo})
   
   if err:=srv.Run(); err!=nil{
	   fmt.Println(err)
   }
}

func createDummyData(repo repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}