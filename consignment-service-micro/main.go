package main

import(
	"context"
	 pb "github.com/avadhut123pisal/GolangMicroservice/consignment-service-micro/proto/consignment"
	 vesselpb "github.com/avadhut123pisal/GolangMicroservice/vessel-service/proto/vessel"
	 micro "github.com/micro/go-micro"
	 microclient "github.com/micro/go-micro/client"
	 "log"
	 "fmt"
)

const (
	port = ":50051"
)

type IRepository interface{
	Create(*pb.Consignment) (*pb.Consignment, error)
    GetAll() []*pb.Consignment 
}

type Repository struct{
	consignments []*pb.Consignment
}

type service struct{
  repo IRepository
  vesselClient vesselpb.VesselServiceClient
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
  repo.consignments= append(repo.consignments, consignment)
  return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

func (s service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) (error) {
	consignments:=s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func (s service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) (error) {
	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err:=s.vesselClient.FindAvailable(context.Background(), &vesselpb.Specification{
	  Capacity: int32(len(req.Containers)),
	  MaxWeight: req.Weight,
	})
	if err!=nil {
		return err
	}
	fmt.Printf("\nfound vessel: %v", vesselResponse.Vessel.Name)
	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id
	consignment, err:=s.repo.Create(req)
	if err!=nil {
		return err
	}
	log.Println("Consignemt created")
	res.Created = true
	res.Consignment = consignment
	return nil
}

func main() {
 repo:=&Repository{}
  // set up gRPC server
 /*  listener, err:=net.Listen("tcp", port)
  if err!=nil{
	  log.Fatalf("failed to listen: %v", err)
  }
  s:=grpc.NewServer()
  //register service on gRPC server
  pb.RegisterShippingServiceServer(s, service{repo: repo})
  // register reflection service on gRPC server
  reflection.Register(s) */
  
  srv :=micro.NewService(
	  micro.Name("go.micro.srv.consignment"),
	  micro.Version("latest"),
  )
  // init will parse command line flags
  srv.Init()
  
  vesselClient:=vesselpb.NewVesselServiceClient("go.micro.srv.vessel", microclient.DefaultClient)
  
  
  // Register handler
  pb.RegisterShippingServiceHandler(srv.Server(), &service{repo: repo, vesselClient: vesselClient})
  
//   log.Printf("Server is running at %v", port)
  if err:=srv.Run(); err!=nil {
	  log.Fatalf("failed to serve: %v", err)
  }
}