// Code generated by goa v3.10.2, DO NOT EDIT.
//
// item gRPC server
//
// Command:
// $ goa gen gaucho/services/item/design -o services/item

package server

import (
	"context"
	"errors"
	itempb "gaucho/services/item/gen/grpc/item/pb"
	item "gaucho/services/item/gen/item"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the itempb.ItemServer interface.
type Server struct {
	GetH    goagrpc.UnaryHandler
	GetAllH goagrpc.UnaryHandler
	AddH    goagrpc.UnaryHandler
	UpdateH goagrpc.UnaryHandler
	RemoveH goagrpc.UnaryHandler
	itempb.UnimplementedItemServer
}

// New instantiates the server struct with the item service endpoints.
func New(e *item.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		GetH:    NewGetHandler(e.Get, uh),
		GetAllH: NewGetAllHandler(e.GetAll, uh),
		AddH:    NewAddHandler(e.Add, uh),
		UpdateH: NewUpdateHandler(e.Update, uh),
		RemoveH: NewRemoveHandler(e.Remove, uh),
	}
}

// NewGetHandler creates a gRPC handler which serves the "item" service "get"
// endpoint.
func NewGetHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetRequest, EncodeGetResponse)
	}
	return h
}

// Get implements the "Get" method in itempb.ItemServer interface.
func (s *Server) Get(ctx context.Context, message *itempb.GetRequest) (*itempb.GetResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "get")
	ctx = context.WithValue(ctx, goa.ServiceKey, "item")
	resp, err := s.GetH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*itempb.GetResponse), nil
}

// NewGetAllHandler creates a gRPC handler which serves the "item" service
// "get_all" endpoint.
func NewGetAllHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, nil, EncodeGetAllResponse)
	}
	return h
}

// GetAll implements the "GetAll" method in itempb.ItemServer interface.
func (s *Server) GetAll(ctx context.Context, message *itempb.GetAllRequest) (*itempb.GetAllResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "get_all")
	ctx = context.WithValue(ctx, goa.ServiceKey, "item")
	resp, err := s.GetAllH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*itempb.GetAllResponse), nil
}

// NewAddHandler creates a gRPC handler which serves the "item" service "add"
// endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in itempb.ItemServer interface.
func (s *Server) Add(ctx context.Context, message *itempb.AddRequest) (*itempb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "item")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "already_exists":
				return nil, goagrpc.NewStatusError(codes.AlreadyExists, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*itempb.AddResponse), nil
}

// NewUpdateHandler creates a gRPC handler which serves the "item" service
// "update" endpoint.
func NewUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateRequest, EncodeUpdateResponse)
	}
	return h
}

// Update implements the "Update" method in itempb.ItemServer interface.
func (s *Server) Update(ctx context.Context, message *itempb.UpdateRequest) (*itempb.UpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "item")
	resp, err := s.UpdateH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "not_found":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			case "already_exists":
				return nil, goagrpc.NewStatusError(codes.AlreadyExists, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*itempb.UpdateResponse), nil
}

// NewRemoveHandler creates a gRPC handler which serves the "item" service
// "remove" endpoint.
func NewRemoveHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRemoveRequest, EncodeRemoveResponse)
	}
	return h
}

// Remove implements the "Remove" method in itempb.ItemServer interface.
func (s *Server) Remove(ctx context.Context, message *itempb.RemoveRequest) (*itempb.RemoveResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "remove")
	ctx = context.WithValue(ctx, goa.ServiceKey, "item")
	resp, err := s.RemoveH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "not_found":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*itempb.RemoveResponse), nil
}
