package main

import (
	pb "github.com/avadhut123pisal/GolangMicroservice/vessel-service/proto/vessel"
	// pb "GolangMicroservice/vessel-service/proto/vessel"
	"errors"
	"context"
)

type handler struct {
  repo *VesselRepository
}

func (h *handler) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	vessel, err:=h.repo.FindAvailable(spec)
	if err!=nil {
		return err
	}
	if vessel == nil {
		return errors.New("No vessel found for specified specifications!")	
	}
	resp.Vessel = vessel
	return nil
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	if err:=h.repo.Create(req); err!=nil{
		return err
	}
	resp.Vessel = req
	resp.Created = true
	return nil	
}