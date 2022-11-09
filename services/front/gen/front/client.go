// Code generated by goa v3.10.2, DO NOT EDIT.
//
// front client
//
// Command:
// $ goa gen gaucho/services/front/design -o services/front

package front

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "front" service client.
type Client struct {
	ShowCharacterEndpoint   goa.Endpoint
	ListCharactersEndpoint  goa.Endpoint
	AddCharacterEndpoint    goa.Endpoint
	UpdateCharacterEndpoint goa.Endpoint
	RemoveCharacterEndpoint goa.Endpoint
	ShowInventoryEndpoint   goa.Endpoint
	UpdateInventoryEndpoint goa.Endpoint
	RemoveInventoryEndpoint goa.Endpoint
	ShowItemEndpoint        goa.Endpoint
	ListItemsEndpoint       goa.Endpoint
	AddItemEndpoint         goa.Endpoint
	UpdateItemEndpoint      goa.Endpoint
	RemoveItemEndpoint      goa.Endpoint
}

// NewClient initializes a "front" service client given the endpoints.
func NewClient(showCharacter, listCharacters, addCharacter, updateCharacter, removeCharacter, showInventory, updateInventory, removeInventory, showItem, listItems, addItem, updateItem, removeItem goa.Endpoint) *Client {
	return &Client{
		ShowCharacterEndpoint:   showCharacter,
		ListCharactersEndpoint:  listCharacters,
		AddCharacterEndpoint:    addCharacter,
		UpdateCharacterEndpoint: updateCharacter,
		RemoveCharacterEndpoint: removeCharacter,
		ShowInventoryEndpoint:   showInventory,
		UpdateInventoryEndpoint: updateInventory,
		RemoveInventoryEndpoint: removeInventory,
		ShowItemEndpoint:        showItem,
		ListItemsEndpoint:       listItems,
		AddItemEndpoint:         addItem,
		UpdateItemEndpoint:      updateItem,
		RemoveItemEndpoint:      removeItem,
	}
}

// ShowCharacter calls the "show_character" endpoint of the "front" service.
// ShowCharacter may return the following errors:
//   - "not_found" (type *goa.ServiceError): Character not found
//   - error: internal error
func (c *Client) ShowCharacter(ctx context.Context, p *ShowCharacterPayload) (res *StoredCharacter, err error) {
	var ires interface{}
	ires, err = c.ShowCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredCharacter), nil
}

// ListCharacters calls the "list_characters" endpoint of the "front" service.
func (c *Client) ListCharacters(ctx context.Context) (res []*StoredCharacter, err error) {
	var ires interface{}
	ires, err = c.ListCharactersEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*StoredCharacter), nil
}

// AddCharacter calls the "add_character" endpoint of the "front" service.
// AddCharacter may return the following errors:
//   - "already_exists" (type *goa.ServiceError): That name already exists
//   - error: internal error
func (c *Client) AddCharacter(ctx context.Context, p *Character) (res *StoredCharacter, err error) {
	var ires interface{}
	ires, err = c.AddCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredCharacter), nil
}

// UpdateCharacter calls the "update_character" endpoint of the "front" service.
// UpdateCharacter may return the following errors:
//   - "not_found" (type *goa.ServiceError): Character not found
//   - "already_exists" (type *goa.ServiceError): That name already exists
//   - error: internal error
func (c *Client) UpdateCharacter(ctx context.Context, p *UpdateCharacterPayload) (res *StoredCharacter, err error) {
	var ires interface{}
	ires, err = c.UpdateCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredCharacter), nil
}

// RemoveCharacter calls the "remove_character" endpoint of the "front" service.
// RemoveCharacter may return the following errors:
//   - "not_found" (type *goa.ServiceError): Character not found
//   - error: internal error
func (c *Client) RemoveCharacter(ctx context.Context, p *RemoveCharacterPayload) (err error) {
	_, err = c.RemoveCharacterEndpoint(ctx, p)
	return
}

// ShowInventory calls the "show_inventory" endpoint of the "front" service.
func (c *Client) ShowInventory(ctx context.Context, p *ShowInventoryPayload) (res []uint32, err error) {
	var ires interface{}
	ires, err = c.ShowInventoryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]uint32), nil
}

// UpdateInventory calls the "update_inventory" endpoint of the "front" service.
func (c *Client) UpdateInventory(ctx context.Context, p *UpdateInventoryPayload) (res []uint32, err error) {
	var ires interface{}
	ires, err = c.UpdateInventoryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]uint32), nil
}

// RemoveInventory calls the "remove_inventory" endpoint of the "front" service.
// RemoveInventory may return the following errors:
//   - "not_found" (type *goa.ServiceError): Character or Item not found
//   - error: internal error
func (c *Client) RemoveInventory(ctx context.Context, p *RemoveInventoryPayload) (err error) {
	_, err = c.RemoveInventoryEndpoint(ctx, p)
	return
}

// ShowItem calls the "show_item" endpoint of the "front" service.
// ShowItem may return the following errors:
//   - "not_found" (type *goa.ServiceError): Item not found
//   - error: internal error
func (c *Client) ShowItem(ctx context.Context, p *ShowItemPayload) (res *StoredItem, err error) {
	var ires interface{}
	ires, err = c.ShowItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredItem), nil
}

// ListItems calls the "list_items" endpoint of the "front" service.
func (c *Client) ListItems(ctx context.Context) (res []*StoredItem, err error) {
	var ires interface{}
	ires, err = c.ListItemsEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*StoredItem), nil
}

// AddItem calls the "add_item" endpoint of the "front" service.
// AddItem may return the following errors:
//   - "already_exists" (type *goa.ServiceError): That name already exists
//   - error: internal error
func (c *Client) AddItem(ctx context.Context, p *Item) (res *StoredItem, err error) {
	var ires interface{}
	ires, err = c.AddItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredItem), nil
}

// UpdateItem calls the "update_item" endpoint of the "front" service.
// UpdateItem may return the following errors:
//   - "not_found" (type *goa.ServiceError): Item not found
//   - "already_exists" (type *goa.ServiceError): That name already exists
//   - error: internal error
func (c *Client) UpdateItem(ctx context.Context, p *UpdateItemPayload) (res *StoredItem, err error) {
	var ires interface{}
	ires, err = c.UpdateItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredItem), nil
}

// RemoveItem calls the "remove_item" endpoint of the "front" service.
// RemoveItem may return the following errors:
//   - "not_found" (type *goa.ServiceError): Item not found
//   - error: internal error
func (c *Client) RemoveItem(ctx context.Context, p *RemoveItemPayload) (err error) {
	_, err = c.RemoveItemEndpoint(ctx, p)
	return
}