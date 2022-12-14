// Code generated by goa v3.10.2, DO NOT EDIT.
//
// inventory gRPC server
//
// Command:
// $ goa gen gaucho/services/inventory/design -o services/inventory

package server

import (
	"context"
	inventorypb "gaucho/services/inventory/gen/grpc/inventory/pb"
	inventory "gaucho/services/inventory/gen/inventory"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the inventorypb.InventoryServer interface.
type Server struct {
	GetH    goagrpc.UnaryHandler
	AddH    goagrpc.UnaryHandler
	RemoveH goagrpc.UnaryHandler
	inventorypb.UnimplementedInventoryServer
}

// New instantiates the server struct with the inventory service endpoints.
func New(e *inventory.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		GetH:    NewGetHandler(e.Get, uh),
		AddH:    NewAddHandler(e.Add, uh),
		RemoveH: NewRemoveHandler(e.Remove, uh),
	}
}

// NewGetHandler creates a gRPC handler which serves the "inventory" service
// "get" endpoint.
func NewGetHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetRequest, EncodeGetResponse)
	}
	return h
}

// Get implements the "Get" method in inventorypb.InventoryServer interface.
func (s *Server) Get(ctx context.Context, message *inventorypb.GetRequest) (*inventorypb.GetResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "get")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.GetH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.GetResponse), nil
}

// NewAddHandler creates a gRPC handler which serves the "inventory" service
// "add" endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in inventorypb.InventoryServer interface.
func (s *Server) Add(ctx context.Context, message *inventorypb.AddRequest) (*inventorypb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.AddResponse), nil
}

// NewRemoveHandler creates a gRPC handler which serves the "inventory" service
// "remove" endpoint.
func NewRemoveHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRemoveRequest, EncodeRemoveResponse)
	}
	return h
}

// Remove implements the "Remove" method in inventorypb.InventoryServer
// interface.
func (s *Server) Remove(ctx context.Context, message *inventorypb.RemoveRequest) (*inventorypb.RemoveResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "remove")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventory")
	resp, err := s.RemoveH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventorypb.RemoveResponse), nil
}
