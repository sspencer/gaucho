// Code generated by goa v3.10.2, DO NOT EDIT.
//
// front service
//
// Command:
// $ goa gen gaucho/services/front/design -o services/front

package front

import (
	"context"
	frontviews "gaucho/services/front/gen/front/views"

	goa "goa.design/goa/v3/pkg"
)

// Public HTTP frontend
type Service interface {
	// Get character by ID
	ShowCharacter(context.Context, *ShowCharacterPayload) (res *StoredCharacter, err error)
	// Get all characters
	ListCharacters(context.Context) (res []*StoredCharacter, err error)
	// Create new character
	AddCharacter(context.Context, *Character) (res *StoredCharacter, err error)
	// Update new character
	UpdateCharacter(context.Context, *UpdateCharacterPayload) (res *StoredCharacter, err error)
	// Remove character from storage
	RemoveCharacter(context.Context, *RemoveCharacterPayload) (err error)
	// Show all items held by a character
	ShowInventory(context.Context, *ShowInventoryPayload) (res []uint32, err error)
	// Add item to character's inventory
	UpdateInventory(context.Context, *UpdateInventoryPayload) (res []uint32, err error)
	// Remove item from character's inventory
	RemoveInventory(context.Context, *RemoveInventoryPayload) (err error)
	// Get item by ID
	ShowItem(context.Context, *ShowItemPayload) (res *StoredItem, err error)
	// Get all items
	ListItems(context.Context) (res []*StoredItem, err error)
	// Create new item
	AddItem(context.Context, *Item) (res *StoredItem, err error)
	// Update new item
	UpdateItem(context.Context, *UpdateItemPayload) (res *StoredItem, err error)
	// Remove item from storage
	RemoveItem(context.Context, *RemoveItemPayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "front"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [13]string{"show_character", "list_characters", "add_character", "update_character", "remove_character", "show_inventory", "update_inventory", "remove_inventory", "show_item", "list_items", "add_item", "update_item", "remove_item"}

// Character is the payload type of the front service add_character method.
type Character struct {
	// Unique Character Name
	Name string
	// Character Description
	Description string
	// Character Health
	Health int
	// Character Experience
	Experience int
}

// Item is the payload type of the front service add_item method.
type Item struct {
	// Unique Item Name
	Name string
	// Item Description
	Description *string
	// Amount of damage item causes
	Damage int
	// Amount of healing item generates
	Healing int
	// Amount of protection item provides
	Protection int
}

// RemoveCharacterPayload is the payload type of the front service
// remove_character method.
type RemoveCharacterPayload struct {
	// ID of character to remove
	ID uint32
}

// RemoveInventoryPayload is the payload type of the front service
// remove_inventory method.
type RemoveInventoryPayload struct {
	// Character ID
	ID uint32
	// Item ID
	ItemID uint32
}

// RemoveItemPayload is the payload type of the front service remove_item
// method.
type RemoveItemPayload struct {
	// ID of item to remove
	ID uint32
}

// ShowCharacterPayload is the payload type of the front service show_character
// method.
type ShowCharacterPayload struct {
	// ID of character to retrieve
	ID uint32
}

// ShowInventoryPayload is the payload type of the front service show_inventory
// method.
type ShowInventoryPayload struct {
	// Character ID
	ID uint32
}

// ShowItemPayload is the payload type of the front service show_item method.
type ShowItemPayload struct {
	// ID of item to retrieve
	ID uint32
}

// StoredCharacter is the result type of the front service show_character
// method.
type StoredCharacter struct {
	// ID is the unique id of the character.
	ID uint32
	// Unique Character Name
	Name string
	// Character Description
	Description string
	// Character Experience
	Experience int
	// Character Health
	Health int
}

// StoredItem is the result type of the front service show_item method.
type StoredItem struct {
	// ID is the unique id of the item.
	ID uint32
	// Unique Item Name
	Name string
	// Item Description
	Description *string
	// Amount of damage item causes
	Damage int
	// Amount of healing item generates
	Healing int
	// Amount of protection item provides
	Protection int
}

// UpdateCharacterPayload is the payload type of the front service
// update_character method.
type UpdateCharacterPayload struct {
	// ID of character to update
	ID uint32
	// Character values to update
	Character *Character
}

// UpdateInventoryPayload is the payload type of the front service
// update_inventory method.
type UpdateInventoryPayload struct {
	// Character ID
	ID uint32
	// Item ID
	ItemID uint32
}

// UpdateItemPayload is the payload type of the front service update_item
// method.
type UpdateItemPayload struct {
	// ID of item to update
	ID uint32
	// Item values to update
	Item *Item
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not_found", false, false, false)
}

// MakeAlreadyExists builds a goa.ServiceError from an error.
func MakeAlreadyExists(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "already_exists", false, false, false)
}

// NewStoredCharacter initializes result type StoredCharacter from viewed
// result type StoredCharacter.
func NewStoredCharacter(vres *frontviews.StoredCharacter) *StoredCharacter {
	return newStoredCharacter(vres.Projected)
}

// NewViewedStoredCharacter initializes viewed result type StoredCharacter from
// result type StoredCharacter using the given view.
func NewViewedStoredCharacter(res *StoredCharacter, view string) *frontviews.StoredCharacter {
	p := newStoredCharacterView(res)
	return &frontviews.StoredCharacter{Projected: p, View: "default"}
}

// NewStoredItem initializes result type StoredItem from viewed result type
// StoredItem.
func NewStoredItem(vres *frontviews.StoredItem) *StoredItem {
	return newStoredItem(vres.Projected)
}

// NewViewedStoredItem initializes viewed result type StoredItem from result
// type StoredItem using the given view.
func NewViewedStoredItem(res *StoredItem, view string) *frontviews.StoredItem {
	p := newStoredItemView(res)
	return &frontviews.StoredItem{Projected: p, View: "default"}
}

// newStoredCharacter converts projected type StoredCharacter to service type
// StoredCharacter.
func newStoredCharacter(vres *frontviews.StoredCharacterView) *StoredCharacter {
	res := &StoredCharacter{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.Health != nil {
		res.Health = *vres.Health
	}
	if vres.Experience != nil {
		res.Experience = *vres.Experience
	}
	return res
}

// newStoredCharacterView projects result type StoredCharacter to projected
// type StoredCharacterView using the "default" view.
func newStoredCharacterView(res *StoredCharacter) *frontviews.StoredCharacterView {
	vres := &frontviews.StoredCharacterView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: &res.Description,
		Experience:  &res.Experience,
		Health:      &res.Health,
	}
	return vres
}

// newStoredItem converts projected type StoredItem to service type StoredItem.
func newStoredItem(vres *frontviews.StoredItemView) *StoredItem {
	res := &StoredItem{
		Description: vres.Description,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Damage != nil {
		res.Damage = *vres.Damage
	}
	if vres.Healing != nil {
		res.Healing = *vres.Healing
	}
	if vres.Protection != nil {
		res.Protection = *vres.Protection
	}
	return res
}

// newStoredItemView projects result type StoredItem to projected type
// StoredItemView using the "default" view.
func newStoredItemView(res *StoredItem) *frontviews.StoredItemView {
	vres := &frontviews.StoredItemView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: res.Description,
		Damage:      &res.Damage,
		Healing:     &res.Healing,
		Protection:  &res.Protection,
	}
	return vres
}