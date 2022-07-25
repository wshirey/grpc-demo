package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/wshirey/grpc-demo/addresses"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// implements AddressServiceServer
type Service struct {
	// UnimplementedAddressServiceServer must be embedded to have forward compatible implementations.
	addresses.UnimplementedAddressesServer
	repo map[string]*addresses.Address
}

func NewService() *Service {
	return &Service{
		repo: make(map[string]*addresses.Address),
	}
}

func (s Service) CreateAddress(ctx context.Context, request *addresses.CreateAddressRequest) (*addresses.Address, error) {
	log.Printf("CreateAddress")
	addr := addresses.Address{
		Id:     fmt.Sprintf("%d", rand.Int31()),
		Street: request.Street,
		City:   request.City,
		Zip:    request.Zip,
	}

	s.repo[addr.Id] = &addr
	return &addr, nil
}

func (s Service) GetAddress(ctx context.Context, request *addresses.GetAddressRequest) (*addresses.Address, error) {
	log.Printf("GetAddress")
	if addr, ok := s.repo[request.Id]; ok {
		return addr, nil
	} else {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}
}

func (s Service) DeleteAddress(ctx context.Context, request *addresses.DeleteAddressRequest) (*addresses.Address, error) {
	log.Printf("DeleteAddress")
	if addr, ok := s.repo[request.Id]; ok {
		delete(s.repo, request.Id)
		return addr, nil
	} else {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}
}
