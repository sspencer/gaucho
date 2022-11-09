// Code generated by goa v3.10.2, DO NOT EDIT.
//
// item service
//
// Command:
// $ goa gen gaucho/services/item/design -o services/item

package item

import (
	"context"
	itemviews "gaucho/services/item/gen/item/views"
)

// Service that provides item API
type Service interface {
	// Retrieve item by the given id
	Get(context.Context, uint32) (res *StoredItem, err error)
	// Retrieve all items
	GetAll(context.Context) (res []*StoredItem, err error)
	// Create an item
	Add(context.Context, *Item) (res *StoredItem, err error)
	// Update an item
	Update(context.Context, *UpdatePayload) (res *StoredItem, err error)
	// Remove item
	Remove(context.Context, *RemovePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "item"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"get", "get_all", "add", "update", "remove"}

// Item is the payload type of the item service add method.
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

// RemovePayload is the payload type of the item service remove method.
type RemovePayload struct {
	// ID of item to remove
	ID uint32
}

// StoredItem is the result type of the item service get method.
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

// UpdatePayload is the payload type of the item service update method.
type UpdatePayload struct {
	// ID of item to update
	ID uint32
	// Item fields to update
	Item *Item
}

// Item name already exists
type AlreadyExists string

// Item not found
type NotFound string

// Error returns an error description.
func (e AlreadyExists) Error() string {
	return "Item name already exists"
}

// ErrorName returns "already_exists".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e AlreadyExists) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "already_exists".
func (e AlreadyExists) GoaErrorName() string {
	return "already_exists"
}

// Error returns an error description.
func (e NotFound) Error() string {
	return "Item not found"
}

// ErrorName returns "not_found".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e NotFound) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "not_found".
func (e NotFound) GoaErrorName() string {
	return "not_found"
}

// NewStoredItem initializes result type StoredItem from viewed result type
// StoredItem.
func NewStoredItem(vres *itemviews.StoredItem) *StoredItem {
	return newStoredItem(vres.Projected)
}

// NewViewedStoredItem initializes viewed result type StoredItem from result
// type StoredItem using the given view.
func NewViewedStoredItem(res *StoredItem, view string) *itemviews.StoredItem {
	p := newStoredItemView(res)
	return &itemviews.StoredItem{Projected: p, View: "default"}
}

// newStoredItem converts projected type StoredItem to service type StoredItem.
func newStoredItem(vres *itemviews.StoredItemView) *StoredItem {
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
func newStoredItemView(res *StoredItem) *itemviews.StoredItemView {
	vres := &itemviews.StoredItemView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: res.Description,
		Damage:      &res.Damage,
		Healing:     &res.Healing,
		Protection:  &res.Protection,
	}
	return vres
}
