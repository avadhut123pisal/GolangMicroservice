package main

import (
	"fmt"
	pb "github.com/avadhut123pisal/GolangProject/GolangMicroservice/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
	"errors"
	"context"
)

type Repository interface {
	FindAvailable(*pb.Specification)  (*pb.Vessel, error)
}

type VesselRepository struct{
  vessels []*pb.Vessel	
}

type service struct {
  	repo Repository
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error){
  for _,vessel:=range repo.vessels {
	  if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
		  return vessel, nil
	  } 
  }
  return nil, errors.New("No vessel found for specified specifications!")	
}

func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	vessel, err:=s.repo.FindAvailable(spec)
	if err!=nil {
		return err
	}
	resp.Vessel = vessel
	return nil
}

func main() {
	// initialse data
   vessels:= []*pb.Vessel{&pb.Vessel{Id: "vessel01", Name: "Medicine vessel", Capacity: 5, MaxWeight: 20000}}
   repo:=&VesselRepository{vessels: vessels}
   srv:=micro.NewService(
	   micro.Name("go.micro.srv.vessel"),
	   micro.Version("latest"),
   )
   
   srv.Init()
   
   pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})
   
   if err:=srv.Run(); err!=nil{
	   fmt.Println(err)
   }
}