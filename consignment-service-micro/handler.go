// consignment-service-micro/handler

package main

import(
	pb "github.com/avadhut123pisal/GolangMicroservice/consignment-service-micro/proto/consignment"
	vesselpb "github.com/avadhut123pisal/GolangMicroservice/vessel-service/proto/vessel"
	
	/* pb "GolangMicroservice/consignment-service-micro/proto/consignment"
	vesselpb "GolangMicroservice/vessel-service/proto/vessel" */
	"gopkg.in/mgo.v2"
	"context"
	"log"
)

type handler struct {
	session      *mgo.Session
	vesselClient vesselpb.VesselServiceClient
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server

func (h handler) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	// create session
	repo :=h.GetRepo()
	defer repo.Close()
  // check for available vessel
  vesselResponse, err:=h.vesselClient.FindAvailable(ctx, &vesselpb.Specification{
	Capacity: int32(len(req.Containers)),
	MaxWeight: req.Weight,  
  })
  if err != nil {
	return err
  }
  log.Printf("found vessel: %s\n", vesselResponse.Vessel.Name)
  // We set the VesselId as the vessel we got back from our
  // vessel service
  req.VesselId =  vesselResponse.Vessel.Id
  // Save our consignment
  if err:=repo.Create(req); err!=nil {
	  return err
  }
  resp.Created = true
  resp.Consignment  = req
  return nil  
}

func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	// create session
	repo :=h.GetRepo()
	defer repo.Close()
	consignments, err:=repo.GetAll()
	if err != nil {
		return err
	}
	resp.Consignments = consignments
	return nil
}

func (h *handler) GetRepo() repository {
	return &ConsignmentRepository{h.session.Clone()}
}

