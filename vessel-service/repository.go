// vessel-service/repository.go

package main

import(
	pb "github.com/avadhut123pisal/GolangMicroservice/vessel-service/proto/vessel"  	
	// pb "GolangMicroservice/vessel-service/proto/vessel" 
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

type repository interface {
	FindAvailable(*pb.Specification)  (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

type VesselRepository struct {
	 session *mgo.Session
}

// Here we're asking for a vessel who's max weight and
// capacity are greater than and equal to the given capacity and weight.
func (repo *VesselRepository) FindAvailable(req *pb.Specification) (*pb.Vessel, error) {
	 var vessel *pb.Vessel
     err:=repo.Collection().Find(bson.M{"capacity": bson.M{"$gte": req.Capacity},
		 "maxweight": bson.M{"$gte": req.MaxWeight},
	 }).One(&vessel)
	 
	 if err != nil {
		return nil, err
	 }
	return vessel, nil
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error{
	err:=repo.Collection().Insert(vessel)
	return err
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) Collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}



