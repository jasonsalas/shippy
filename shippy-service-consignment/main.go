package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/jasonsalas/shippy/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository is a dummy datastore for now until a more robust datastore is implemented
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create satisfies the local 'repository' interface
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()

	return consignment, nil
}

// GetAll saisfies the local 'repository' interface
// and return all consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// service satisfies the service in the protobuf definition
type service struct {
	repo repository
}

// CreateConsignment is handled by the gRPC server
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// save the consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// return matches the 'Response' message in the protobuf definition
	/*
			message Response {
			bool created = 1;
			Consignment consignment = 2;
		}
	*/
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

// GetConsignments is handled by the gRPC server
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()

	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &Repository{}

	// instantiate a gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listed: %v", err)
	}
	s := grpc.NewServer()

	// register the service with the gRPC server
	pb.RegisterShippingServiceServer(s, &service{repo})
	reflection.Register(s)

	log.Println("running on port: ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
