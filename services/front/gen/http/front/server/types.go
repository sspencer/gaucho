// Code generated by goa v3.10.2, DO NOT EDIT.
//
// front HTTP server types
//
// Command:
// $ goa gen gaucho/services/front/design -o services/front

package server

import (
	front "gaucho/services/front/gen/front"
	frontviews "gaucho/services/front/gen/front/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// AddCharacterRequestBody is the type of the "front" service "add_character"
// endpoint HTTP request body.
type AddCharacterRequestBody struct {
	// Unique Character Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Character Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Character Health
	Health *int `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Character Experience
	Experience *int `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// UpdateCharacterRequestBody is the type of the "front" service
// "update_character" endpoint HTTP request body.
type UpdateCharacterRequestBody struct {
	// Unique Character Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Character Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Character Health
	Health *int `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Character Experience
	Experience *int `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// AddItemRequestBody is the type of the "front" service "add_item" endpoint
// HTTP request body.
type AddItemRequestBody struct {
	// Unique Item Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Item Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of damage item causes
	Damage *int `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Amount of healing item generates
	Healing *int `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Amount of protection item provides
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// UpdateItemRequestBody is the type of the "front" service "update_item"
// endpoint HTTP request body.
type UpdateItemRequestBody struct {
	// Unique Item Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Item Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of damage item causes
	Damage *int `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Amount of healing item generates
	Healing *int `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Amount of protection item provides
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// ShowCharacterResponseBody is the type of the "front" service
// "show_character" endpoint HTTP response body.
type ShowCharacterResponseBody struct {
	// ID is the unique id of the character.
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Unique Character Name
	Name string `form:"name" json:"name" xml:"name"`
	// Character Description
	Description string `form:"description" json:"description" xml:"description"`
	// Character Health
	Health int `form:"health" json:"health" xml:"health"`
	// Character Experience
	Experience int `form:"experience" json:"experience" xml:"experience"`
}

// ListCharactersResponseBody is the type of the "front" service
// "list_characters" endpoint HTTP response body.
type ListCharactersResponseBody []*StoredCharacterResponse

// AddCharacterResponseBody is the type of the "front" service "add_character"
// endpoint HTTP response body.
type AddCharacterResponseBody struct {
	// ID is the unique id of the character.
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Unique Character Name
	Name string `form:"name" json:"name" xml:"name"`
	// Character Description
	Description string `form:"description" json:"description" xml:"description"`
	// Character Health
	Health int `form:"health" json:"health" xml:"health"`
	// Character Experience
	Experience int `form:"experience" json:"experience" xml:"experience"`
}

// UpdateCharacterResponseBody is the type of the "front" service
// "update_character" endpoint HTTP response body.
type UpdateCharacterResponseBody struct {
	// ID is the unique id of the character.
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Unique Character Name
	Name string `form:"name" json:"name" xml:"name"`
	// Character Description
	Description string `form:"description" json:"description" xml:"description"`
	// Character Health
	Health int `form:"health" json:"health" xml:"health"`
	// Character Experience
	Experience int `form:"experience" json:"experience" xml:"experience"`
}

// ShowItemResponseBody is the type of the "front" service "show_item" endpoint
// HTTP response body.
type ShowItemResponseBody struct {
	// ID is the unique id of the item.
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Unique Item Name
	Name string `form:"name" json:"name" xml:"name"`
	// Item Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of damage item causes
	Damage int `form:"damage" json:"damage" xml:"damage"`
	// Amount of healing item generates
	Healing int `form:"healing" json:"healing" xml:"healing"`
	// Amount of protection item provides
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// ListItemsResponseBody is the type of the "front" service "list_items"
// endpoint HTTP response body.
type ListItemsResponseBody []*StoredItemResponse

// AddItemResponseBody is the type of the "front" service "add_item" endpoint
// HTTP response body.
type AddItemResponseBody struct {
	// ID is the unique id of the item.
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Unique Item Name
	Name string `form:"name" json:"name" xml:"name"`
	// Item Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of damage item causes
	Damage int `form:"damage" json:"damage" xml:"damage"`
	// Amount of healing item generates
	Healing int `form:"healing" json:"healing" xml:"healing"`
	// Amount of protection item provides
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// UpdateItemResponseBody is the type of the "front" service "update_item"
// endpoint HTTP response body.
type UpdateItemResponseBody struct {
	// ID is the unique id of the item.
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Unique Item Name
	Name string `form:"name" json:"name" xml:"name"`
	// Item Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of damage item causes
	Damage int `form:"damage" json:"damage" xml:"damage"`
	// Amount of healing item generates
	Healing int `form:"healing" json:"healing" xml:"healing"`
	// Amount of protection item provides
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// ShowCharacterNotFoundResponseBody is the type of the "front" service
// "show_character" endpoint HTTP response body for the "not_found" error.
type ShowCharacterNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// AddCharacterAlreadyExistsResponseBody is the type of the "front" service
// "add_character" endpoint HTTP response body for the "already_exists" error.
type AddCharacterAlreadyExistsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateCharacterNotFoundResponseBody is the type of the "front" service
// "update_character" endpoint HTTP response body for the "not_found" error.
type UpdateCharacterNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateCharacterAlreadyExistsResponseBody is the type of the "front" service
// "update_character" endpoint HTTP response body for the "already_exists"
// error.
type UpdateCharacterAlreadyExistsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RemoveCharacterNotFoundResponseBody is the type of the "front" service
// "remove_character" endpoint HTTP response body for the "not_found" error.
type RemoveCharacterNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RemoveInventoryNotFoundResponseBody is the type of the "front" service
// "remove_inventory" endpoint HTTP response body for the "not_found" error.
type RemoveInventoryNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ShowItemNotFoundResponseBody is the type of the "front" service "show_item"
// endpoint HTTP response body for the "not_found" error.
type ShowItemNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// AddItemAlreadyExistsResponseBody is the type of the "front" service
// "add_item" endpoint HTTP response body for the "already_exists" error.
type AddItemAlreadyExistsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateItemNotFoundResponseBody is the type of the "front" service
// "update_item" endpoint HTTP response body for the "not_found" error.
type UpdateItemNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateItemAlreadyExistsResponseBody is the type of the "front" service
// "update_item" endpoint HTTP response body for the "already_exists" error.
type UpdateItemAlreadyExistsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RemoveItemNotFoundResponseBody is the type of the "front" service
// "remove_item" endpoint HTTP response body for the "not_found" error.
type RemoveItemNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// StoredCharacterResponse is used to define fields on response body types.
type StoredCharacterResponse struct {
	// ID is the unique id of the character.
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Unique Character Name
	Name string `form:"name" json:"name" xml:"name"`
	// Character Description
	Description string `form:"description" json:"description" xml:"description"`
	// Character Experience
	Experience int `form:"experience" json:"experience" xml:"experience"`
	// Character Health
	Health int `form:"health" json:"health" xml:"health"`
}

// StoredItemResponse is used to define fields on response body types.
type StoredItemResponse struct {
	// ID is the unique id of the item.
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Unique Item Name
	Name string `form:"name" json:"name" xml:"name"`
	// Item Description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of damage item causes
	Damage int `form:"damage" json:"damage" xml:"damage"`
	// Amount of healing item generates
	Healing int `form:"healing" json:"healing" xml:"healing"`
	// Amount of protection item provides
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// NewShowCharacterResponseBody builds the HTTP response body from the result
// of the "show_character" endpoint of the "front" service.
func NewShowCharacterResponseBody(res *frontviews.StoredCharacterView) *ShowCharacterResponseBody {
	body := &ShowCharacterResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: *res.Description,
		Experience:  *res.Experience,
		Health:      *res.Health,
	}
	return body
}

// NewListCharactersResponseBody builds the HTTP response body from the result
// of the "list_characters" endpoint of the "front" service.
func NewListCharactersResponseBody(res []*front.StoredCharacter) ListCharactersResponseBody {
	body := make([]*StoredCharacterResponse, len(res))
	for i, val := range res {
		body[i] = marshalFrontStoredCharacterToStoredCharacterResponse(val)
	}
	return body
}

// NewAddCharacterResponseBody builds the HTTP response body from the result of
// the "add_character" endpoint of the "front" service.
func NewAddCharacterResponseBody(res *frontviews.StoredCharacterView) *AddCharacterResponseBody {
	body := &AddCharacterResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: *res.Description,
		Experience:  *res.Experience,
		Health:      *res.Health,
	}
	return body
}

// NewUpdateCharacterResponseBody builds the HTTP response body from the result
// of the "update_character" endpoint of the "front" service.
func NewUpdateCharacterResponseBody(res *frontviews.StoredCharacterView) *UpdateCharacterResponseBody {
	body := &UpdateCharacterResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: *res.Description,
		Experience:  *res.Experience,
		Health:      *res.Health,
	}
	return body
}

// NewShowItemResponseBody builds the HTTP response body from the result of the
// "show_item" endpoint of the "front" service.
func NewShowItemResponseBody(res *frontviews.StoredItemView) *ShowItemResponseBody {
	body := &ShowItemResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: res.Description,
		Damage:      *res.Damage,
		Healing:     *res.Healing,
		Protection:  *res.Protection,
	}
	return body
}

// NewListItemsResponseBody builds the HTTP response body from the result of
// the "list_items" endpoint of the "front" service.
func NewListItemsResponseBody(res []*front.StoredItem) ListItemsResponseBody {
	body := make([]*StoredItemResponse, len(res))
	for i, val := range res {
		body[i] = marshalFrontStoredItemToStoredItemResponse(val)
	}
	return body
}

// NewAddItemResponseBody builds the HTTP response body from the result of the
// "add_item" endpoint of the "front" service.
func NewAddItemResponseBody(res *frontviews.StoredItemView) *AddItemResponseBody {
	body := &AddItemResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: res.Description,
		Damage:      *res.Damage,
		Healing:     *res.Healing,
		Protection:  *res.Protection,
	}
	return body
}

// NewUpdateItemResponseBody builds the HTTP response body from the result of
// the "update_item" endpoint of the "front" service.
func NewUpdateItemResponseBody(res *frontviews.StoredItemView) *UpdateItemResponseBody {
	body := &UpdateItemResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: res.Description,
		Damage:      *res.Damage,
		Healing:     *res.Healing,
		Protection:  *res.Protection,
	}
	return body
}

// NewShowCharacterNotFoundResponseBody builds the HTTP response body from the
// result of the "show_character" endpoint of the "front" service.
func NewShowCharacterNotFoundResponseBody(res *goa.ServiceError) *ShowCharacterNotFoundResponseBody {
	body := &ShowCharacterNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewAddCharacterAlreadyExistsResponseBody builds the HTTP response body from
// the result of the "add_character" endpoint of the "front" service.
func NewAddCharacterAlreadyExistsResponseBody(res *goa.ServiceError) *AddCharacterAlreadyExistsResponseBody {
	body := &AddCharacterAlreadyExistsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateCharacterNotFoundResponseBody builds the HTTP response body from
// the result of the "update_character" endpoint of the "front" service.
func NewUpdateCharacterNotFoundResponseBody(res *goa.ServiceError) *UpdateCharacterNotFoundResponseBody {
	body := &UpdateCharacterNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateCharacterAlreadyExistsResponseBody builds the HTTP response body
// from the result of the "update_character" endpoint of the "front" service.
func NewUpdateCharacterAlreadyExistsResponseBody(res *goa.ServiceError) *UpdateCharacterAlreadyExistsResponseBody {
	body := &UpdateCharacterAlreadyExistsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRemoveCharacterNotFoundResponseBody builds the HTTP response body from
// the result of the "remove_character" endpoint of the "front" service.
func NewRemoveCharacterNotFoundResponseBody(res *goa.ServiceError) *RemoveCharacterNotFoundResponseBody {
	body := &RemoveCharacterNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRemoveInventoryNotFoundResponseBody builds the HTTP response body from
// the result of the "remove_inventory" endpoint of the "front" service.
func NewRemoveInventoryNotFoundResponseBody(res *goa.ServiceError) *RemoveInventoryNotFoundResponseBody {
	body := &RemoveInventoryNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewShowItemNotFoundResponseBody builds the HTTP response body from the
// result of the "show_item" endpoint of the "front" service.
func NewShowItemNotFoundResponseBody(res *goa.ServiceError) *ShowItemNotFoundResponseBody {
	body := &ShowItemNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewAddItemAlreadyExistsResponseBody builds the HTTP response body from the
// result of the "add_item" endpoint of the "front" service.
func NewAddItemAlreadyExistsResponseBody(res *goa.ServiceError) *AddItemAlreadyExistsResponseBody {
	body := &AddItemAlreadyExistsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateItemNotFoundResponseBody builds the HTTP response body from the
// result of the "update_item" endpoint of the "front" service.
func NewUpdateItemNotFoundResponseBody(res *goa.ServiceError) *UpdateItemNotFoundResponseBody {
	body := &UpdateItemNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateItemAlreadyExistsResponseBody builds the HTTP response body from
// the result of the "update_item" endpoint of the "front" service.
func NewUpdateItemAlreadyExistsResponseBody(res *goa.ServiceError) *UpdateItemAlreadyExistsResponseBody {
	body := &UpdateItemAlreadyExistsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRemoveItemNotFoundResponseBody builds the HTTP response body from the
// result of the "remove_item" endpoint of the "front" service.
func NewRemoveItemNotFoundResponseBody(res *goa.ServiceError) *RemoveItemNotFoundResponseBody {
	body := &RemoveItemNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewShowCharacterPayload builds a front service show_character endpoint
// payload.
func NewShowCharacterPayload(id uint32) *front.ShowCharacterPayload {
	v := &front.ShowCharacterPayload{}
	v.ID = id

	return v
}

// NewAddCharacterCharacter builds a front service add_character endpoint
// payload.
func NewAddCharacterCharacter(body *AddCharacterRequestBody) *front.Character {
	v := &front.Character{
		Name:        *body.Name,
		Description: *body.Description,
		Health:      *body.Health,
		Experience:  *body.Experience,
	}

	return v
}

// NewUpdateCharacterPayload builds a front service update_character endpoint
// payload.
func NewUpdateCharacterPayload(body *UpdateCharacterRequestBody, id uint32) *front.UpdateCharacterPayload {
	v := &front.Character{
		Name:        *body.Name,
		Description: *body.Description,
		Health:      *body.Health,
		Experience:  *body.Experience,
	}
	res := &front.UpdateCharacterPayload{
		Character: v,
	}
	res.ID = id

	return res
}

// NewRemoveCharacterPayload builds a front service remove_character endpoint
// payload.
func NewRemoveCharacterPayload(id uint32) *front.RemoveCharacterPayload {
	v := &front.RemoveCharacterPayload{}
	v.ID = id

	return v
}

// NewShowInventoryPayload builds a front service show_inventory endpoint
// payload.
func NewShowInventoryPayload(id uint32) *front.ShowInventoryPayload {
	v := &front.ShowInventoryPayload{}
	v.ID = id

	return v
}

// NewUpdateInventoryPayload builds a front service update_inventory endpoint
// payload.
func NewUpdateInventoryPayload(id uint32, itemID uint32) *front.UpdateInventoryPayload {
	v := &front.UpdateInventoryPayload{}
	v.ID = id
	v.ItemID = itemID

	return v
}

// NewRemoveInventoryPayload builds a front service remove_inventory endpoint
// payload.
func NewRemoveInventoryPayload(id uint32, itemID uint32) *front.RemoveInventoryPayload {
	v := &front.RemoveInventoryPayload{}
	v.ID = id
	v.ItemID = itemID

	return v
}

// NewShowItemPayload builds a front service show_item endpoint payload.
func NewShowItemPayload(id uint32) *front.ShowItemPayload {
	v := &front.ShowItemPayload{}
	v.ID = id

	return v
}

// NewAddItemItem builds a front service add_item endpoint payload.
func NewAddItemItem(body *AddItemRequestBody) *front.Item {
	v := &front.Item{
		Name:        *body.Name,
		Description: body.Description,
		Damage:      *body.Damage,
		Healing:     *body.Healing,
		Protection:  *body.Protection,
	}

	return v
}

// NewUpdateItemPayload builds a front service update_item endpoint payload.
func NewUpdateItemPayload(body *UpdateItemRequestBody, id uint32) *front.UpdateItemPayload {
	v := &front.Item{
		Name:        *body.Name,
		Description: body.Description,
		Damage:      *body.Damage,
		Healing:     *body.Healing,
		Protection:  *body.Protection,
	}
	res := &front.UpdateItemPayload{
		Item: v,
	}
	res.ID = id

	return res
}

// NewRemoveItemPayload builds a front service remove_item endpoint payload.
func NewRemoveItemPayload(id uint32) *front.RemoveItemPayload {
	v := &front.RemoveItemPayload{}
	v.ID = id

	return v
}

// ValidateAddCharacterRequestBody runs the validations defined on
// add_character_request_body
func ValidateAddCharacterRequestBody(body *AddCharacterRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "body"))
	}
	if body.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 2, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 50, false))
		}
	}
	if body.Health != nil {
		if *body.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", *body.Health, 0, true))
		}
	}
	if body.Health != nil {
		if *body.Health > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", *body.Health, 100, false))
		}
	}
	if body.Experience != nil {
		if *body.Experience < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", *body.Experience, 1, true))
		}
	}
	if body.Experience != nil {
		if *body.Experience > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", *body.Experience, 100, false))
		}
	}
	return
}

// ValidateUpdateCharacterRequestBody runs the validations defined on
// update_character_request_body
func ValidateUpdateCharacterRequestBody(body *UpdateCharacterRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "body"))
	}
	if body.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 2, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 50, false))
		}
	}
	if body.Health != nil {
		if *body.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", *body.Health, 0, true))
		}
	}
	if body.Health != nil {
		if *body.Health > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", *body.Health, 100, false))
		}
	}
	if body.Experience != nil {
		if *body.Experience < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", *body.Experience, 1, true))
		}
	}
	if body.Experience != nil {
		if *body.Experience > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", *body.Experience, 100, false))
		}
	}
	return
}

// ValidateAddItemRequestBody runs the validations defined on
// add_item_request_body
func ValidateAddItemRequestBody(body *AddItemRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "body"))
	}
	if body.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "body"))
	}
	if body.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 2, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 50, false))
		}
	}
	if body.Damage != nil {
		if *body.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", *body.Damage, 0, true))
		}
	}
	if body.Damage != nil {
		if *body.Damage > 25 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", *body.Damage, 25, false))
		}
	}
	if body.Healing != nil {
		if *body.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", *body.Healing, 0, true))
		}
	}
	if body.Healing != nil {
		if *body.Healing > 50 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", *body.Healing, 50, false))
		}
	}
	if body.Protection != nil {
		if *body.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", *body.Protection, 0, true))
		}
	}
	if body.Protection != nil {
		if *body.Protection > 20 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", *body.Protection, 20, false))
		}
	}
	return
}

// ValidateUpdateItemRequestBody runs the validations defined on
// update_item_request_body
func ValidateUpdateItemRequestBody(body *UpdateItemRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "body"))
	}
	if body.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "body"))
	}
	if body.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 2, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 50, false))
		}
	}
	if body.Damage != nil {
		if *body.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", *body.Damage, 0, true))
		}
	}
	if body.Damage != nil {
		if *body.Damage > 25 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", *body.Damage, 25, false))
		}
	}
	if body.Healing != nil {
		if *body.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", *body.Healing, 0, true))
		}
	}
	if body.Healing != nil {
		if *body.Healing > 50 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", *body.Healing, 50, false))
		}
	}
	if body.Protection != nil {
		if *body.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", *body.Protection, 0, true))
		}
	}
	if body.Protection != nil {
		if *body.Protection > 20 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", *body.Protection, 20, false))
		}
	}
	return
}
