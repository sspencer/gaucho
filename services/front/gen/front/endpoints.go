// Code generated by goa v3.10.2, DO NOT EDIT.
//
// front endpoints
//
// Command:
// $ goa gen gaucho/services/front/design -o services/front

package front

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "front" service endpoints.
type Endpoints struct {
	ShowCharacter   goa.Endpoint
	ListCharacters  goa.Endpoint
	AddCharacter    goa.Endpoint
	UpdateCharacter goa.Endpoint
	RemoveCharacter goa.Endpoint
	ShowInventory   goa.Endpoint
	UpdateInventory goa.Endpoint
	RemoveInventory goa.Endpoint
	ShowItem        goa.Endpoint
	ListItems       goa.Endpoint
	AddItem         goa.Endpoint
	UpdateItem      goa.Endpoint
	RemoveItem      goa.Endpoint
}

// NewEndpoints wraps the methods of the "front" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		ShowCharacter:   NewShowCharacterEndpoint(s),
		ListCharacters:  NewListCharactersEndpoint(s),
		AddCharacter:    NewAddCharacterEndpoint(s),
		UpdateCharacter: NewUpdateCharacterEndpoint(s),
		RemoveCharacter: NewRemoveCharacterEndpoint(s),
		ShowInventory:   NewShowInventoryEndpoint(s),
		UpdateInventory: NewUpdateInventoryEndpoint(s),
		RemoveInventory: NewRemoveInventoryEndpoint(s),
		ShowItem:        NewShowItemEndpoint(s),
		ListItems:       NewListItemsEndpoint(s),
		AddItem:         NewAddItemEndpoint(s),
		UpdateItem:      NewUpdateItemEndpoint(s),
		RemoveItem:      NewRemoveItemEndpoint(s),
	}
}

// Use applies the given middleware to all the "front" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ShowCharacter = m(e.ShowCharacter)
	e.ListCharacters = m(e.ListCharacters)
	e.AddCharacter = m(e.AddCharacter)
	e.UpdateCharacter = m(e.UpdateCharacter)
	e.RemoveCharacter = m(e.RemoveCharacter)
	e.ShowInventory = m(e.ShowInventory)
	e.UpdateInventory = m(e.UpdateInventory)
	e.RemoveInventory = m(e.RemoveInventory)
	e.ShowItem = m(e.ShowItem)
	e.ListItems = m(e.ListItems)
	e.AddItem = m(e.AddItem)
	e.UpdateItem = m(e.UpdateItem)
	e.RemoveItem = m(e.RemoveItem)
}

// NewShowCharacterEndpoint returns an endpoint function that calls the method
// "show_character" of service "front".
func NewShowCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowCharacterPayload)
		res, err := s.ShowCharacter(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredCharacter(res, "default")
		return vres, nil
	}
}

// NewListCharactersEndpoint returns an endpoint function that calls the method
// "list_characters" of service "front".
func NewListCharactersEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListCharacters(ctx)
	}
}

// NewAddCharacterEndpoint returns an endpoint function that calls the method
// "add_character" of service "front".
func NewAddCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Character)
		res, err := s.AddCharacter(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredCharacter(res, "default")
		return vres, nil
	}
}

// NewUpdateCharacterEndpoint returns an endpoint function that calls the
// method "update_character" of service "front".
func NewUpdateCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateCharacterPayload)
		res, err := s.UpdateCharacter(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredCharacter(res, "default")
		return vres, nil
	}
}

// NewRemoveCharacterEndpoint returns an endpoint function that calls the
// method "remove_character" of service "front".
func NewRemoveCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RemoveCharacterPayload)
		return nil, s.RemoveCharacter(ctx, p)
	}
}

// NewShowInventoryEndpoint returns an endpoint function that calls the method
// "show_inventory" of service "front".
func NewShowInventoryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowInventoryPayload)
		return s.ShowInventory(ctx, p)
	}
}

// NewUpdateInventoryEndpoint returns an endpoint function that calls the
// method "update_inventory" of service "front".
func NewUpdateInventoryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateInventoryPayload)
		return s.UpdateInventory(ctx, p)
	}
}

// NewRemoveInventoryEndpoint returns an endpoint function that calls the
// method "remove_inventory" of service "front".
func NewRemoveInventoryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RemoveInventoryPayload)
		return nil, s.RemoveInventory(ctx, p)
	}
}

// NewShowItemEndpoint returns an endpoint function that calls the method
// "show_item" of service "front".
func NewShowItemEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowItemPayload)
		res, err := s.ShowItem(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredItem(res, "default")
		return vres, nil
	}
}

// NewListItemsEndpoint returns an endpoint function that calls the method
// "list_items" of service "front".
func NewListItemsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListItems(ctx)
	}
}

// NewAddItemEndpoint returns an endpoint function that calls the method
// "add_item" of service "front".
func NewAddItemEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Item)
		res, err := s.AddItem(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredItem(res, "default")
		return vres, nil
	}
}

// NewUpdateItemEndpoint returns an endpoint function that calls the method
// "update_item" of service "front".
func NewUpdateItemEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateItemPayload)
		res, err := s.UpdateItem(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredItem(res, "default")
		return vres, nil
	}
}

// NewRemoveItemEndpoint returns an endpoint function that calls the method
// "remove_item" of service "front".
func NewRemoveItemEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RemoveItemPayload)
		return nil, s.RemoveItem(ctx, p)
	}
}
