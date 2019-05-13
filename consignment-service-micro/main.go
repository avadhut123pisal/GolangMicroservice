package main

import(
   pb "github.com/avadhut123pisal/GolangMicroservice/consignment-service-micro/proto/consignment"
	vesselpb "github.com/avadhut123pisal/GolangMicroservice/vessel-service/proto/vessel"
	
	// pb "GolangMicroservice/consignment-service-micro/proto/consignment"
	// vesselpb "GolangMicroservice/vessel-service/proto/vessel"
	 micro "github.com/micro/go-micro"
	 microclient "github.com/micro/go-micro/client"
   "log"
   "os"
)

const (
	DEFAULT_DB_HOST = "localhost:27017"
)

func main() {
  srv :=micro.NewService(
	  micro.Name("go.micro.srv.consignment"),
	  micro.Version("latest"),
  )
  // init will parse command line flags
  srv.Init()
  
  uri:=os.Getenv("DB_HOST")
  if uri == "" {
    uri = DEFAULT_DB_HOST
  }
  session, err:=CreateSession(uri)
  // Mgo creates a 'master' session, we need to end that session
  // before the main function closes.
  defer session.Close()
  if err != nil {
	log.Panicf("Could not connect to datastore with host %s - %v", uri, err)
  }
  
  vesselClient:=vesselpb.NewVesselServiceClient("go.micro.srv.vessel", microclient.DefaultClient)
  
  // Register handler
  pb.RegisterShippingServiceHandler(srv.Server(), &handler{session: session, vesselClient: vesselClient})
  
  if err:=srv.Run(); err!=nil {
	  log.Fatalf("failed to serve: %v", err)
  }
}