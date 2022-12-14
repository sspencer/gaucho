// Code generated by goa v3.10.2, DO NOT EDIT.
//
// character client
//
// Command:
// $ goa gen gaucho/services/character/design -o services/character

package character

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "character" service client.
type Client struct {
	GetEndpoint    goa.Endpoint
	GetAllEndpoint goa.Endpoint
	AddEndpoint    goa.Endpoint
	UpdateEndpoint goa.Endpoint
	RemoveEndpoint goa.Endpoint
}

// NewClient initializes a "character" service client given the endpoints.
func NewClient(get, getAll, add, update, remove goa.Endpoint) *Client {
	return &Client{
		GetEndpoint:    get,
		GetAllEndpoint: getAll,
		AddEndpoint:    add,
		UpdateEndpoint: update,
		RemoveEndpoint: remove,
	}
}

// Get calls the "get" endpoint of the "character" service.
// Get may return the following errors:
//   - "not_found" (type NotFound)
//   - error: internal error
func (c *Client) Get(ctx context.Context, p uint32) (res *StoredCharacter, err error) {
	var ires interface{}
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredCharacter), nil
}

// GetAll calls the "get_all" endpoint of the "character" service.
func (c *Client) GetAll(ctx context.Context) (res []*StoredCharacter, err error) {
	var ires interface{}
	ires, err = c.GetAllEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*StoredCharacter), nil
}

// Add calls the "add" endpoint of the "character" service.
// Add may return the following errors:
//   - "already_exists" (type AlreadyExists)
//   - error: internal error
func (c *Client) Add(ctx context.Context, p *Character) (res *StoredCharacter, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredCharacter), nil
}

// Update calls the "update" endpoint of the "character" service.
// Update may return the following errors:
//   - "already_exists" (type AlreadyExists)
//   - "not_found" (type NotFound)
//   - error: internal error
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (res *StoredCharacter, err error) {
	var ires interface{}
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredCharacter), nil
}

// Remove calls the "remove" endpoint of the "character" service.
// Remove may return the following errors:
//   - "not_found" (type NotFound)
//   - error: internal error
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}
