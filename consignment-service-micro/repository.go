package main

import(
	 pb "github.com/avadhut123pisal/GolangMicroservice/consignment-service-micro/proto/consignment"
	//  pb "GolangMicroservice/consignment-service-micro/proto/consignment"
	 "gopkg.in/mgo.v2"
)

const (
	dbName = "shippy"
	consignmentCollection = "consignments"
)

type repository interface{
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error) 
	Close()
}

type ConsignmentRepository struct{
	session *mgo.Session
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
  err:= repo.Collection().Insert(consignment)
  return err
}

func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
  err:=repo.Collection().Find(nil).All(&consignments)
	return consignments,err
}

func (repo *ConsignmentRepository) Collection() *mgo.Collection{
	return repo.session.DB(dbName).C(consignmentCollection)
}

// Close closes the database session after each query has ran.
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}