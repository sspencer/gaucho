// Code generated by goa v3.10.2, DO NOT EDIT.
//
// front HTTP server
//
// Command:
// $ goa gen gaucho/services/front/design -o services/front

package server

import (
	"context"
	front "gaucho/services/front/gen/front"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the front service endpoint HTTP handlers.
type Server struct {
	Mounts          []*MountPoint
	ShowCharacter   http.Handler
	ListCharacters  http.Handler
	AddCharacter    http.Handler
	UpdateCharacter http.Handler
	RemoveCharacter http.Handler
	ShowInventory   http.Handler
	UpdateInventory http.Handler
	RemoveInventory http.Handler
	ShowItem        http.Handler
	ListItems       http.Handler
	AddItem         http.Handler
	UpdateItem      http.Handler
	RemoveItem      http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the front service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *front.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ShowCharacter", "GET", "/character/{id}"},
			{"ListCharacters", "GET", "/character"},
			{"AddCharacter", "POST", "/character"},
			{"UpdateCharacter", "PUT", "/character/{id}"},
			{"RemoveCharacter", "DELETE", "/character/{id}"},
			{"ShowInventory", "GET", "/inventory/{id}"},
			{"UpdateInventory", "PUT", "/inventory/{id}/item/{item_id}"},
			{"RemoveInventory", "DELETE", "/inventory/{id}/item/{item_id}"},
			{"ShowItem", "GET", "/item/{id}"},
			{"ListItems", "GET", "/item"},
			{"AddItem", "POST", "/item"},
			{"UpdateItem", "PUT", "/item/{id}"},
			{"RemoveItem", "DELETE", "/item/{id}"},
		},
		ShowCharacter:   NewShowCharacterHandler(e.ShowCharacter, mux, decoder, encoder, errhandler, formatter),
		ListCharacters:  NewListCharactersHandler(e.ListCharacters, mux, decoder, encoder, errhandler, formatter),
		AddCharacter:    NewAddCharacterHandler(e.AddCharacter, mux, decoder, encoder, errhandler, formatter),
		UpdateCharacter: NewUpdateCharacterHandler(e.UpdateCharacter, mux, decoder, encoder, errhandler, formatter),
		RemoveCharacter: NewRemoveCharacterHandler(e.RemoveCharacter, mux, decoder, encoder, errhandler, formatter),
		ShowInventory:   NewShowInventoryHandler(e.ShowInventory, mux, decoder, encoder, errhandler, formatter),
		UpdateInventory: NewUpdateInventoryHandler(e.UpdateInventory, mux, decoder, encoder, errhandler, formatter),
		RemoveInventory: NewRemoveInventoryHandler(e.RemoveInventory, mux, decoder, encoder, errhandler, formatter),
		ShowItem:        NewShowItemHandler(e.ShowItem, mux, decoder, encoder, errhandler, formatter),
		ListItems:       NewListItemsHandler(e.ListItems, mux, decoder, encoder, errhandler, formatter),
		AddItem:         NewAddItemHandler(e.AddItem, mux, decoder, encoder, errhandler, formatter),
		UpdateItem:      NewUpdateItemHandler(e.UpdateItem, mux, decoder, encoder, errhandler, formatter),
		RemoveItem:      NewRemoveItemHandler(e.RemoveItem, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "front" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ShowCharacter = m(s.ShowCharacter)
	s.ListCharacters = m(s.ListCharacters)
	s.AddCharacter = m(s.AddCharacter)
	s.UpdateCharacter = m(s.UpdateCharacter)
	s.RemoveCharacter = m(s.RemoveCharacter)
	s.ShowInventory = m(s.ShowInventory)
	s.UpdateInventory = m(s.UpdateInventory)
	s.RemoveInventory = m(s.RemoveInventory)
	s.ShowItem = m(s.ShowItem)
	s.ListItems = m(s.ListItems)
	s.AddItem = m(s.AddItem)
	s.UpdateItem = m(s.UpdateItem)
	s.RemoveItem = m(s.RemoveItem)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return front.MethodNames[:] }

// Mount configures the mux to serve the front endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountShowCharacterHandler(mux, h.ShowCharacter)
	MountListCharactersHandler(mux, h.ListCharacters)
	MountAddCharacterHandler(mux, h.AddCharacter)
	MountUpdateCharacterHandler(mux, h.UpdateCharacter)
	MountRemoveCharacterHandler(mux, h.RemoveCharacter)
	MountShowInventoryHandler(mux, h.ShowInventory)
	MountUpdateInventoryHandler(mux, h.UpdateInventory)
	MountRemoveInventoryHandler(mux, h.RemoveInventory)
	MountShowItemHandler(mux, h.ShowItem)
	MountListItemsHandler(mux, h.ListItems)
	MountAddItemHandler(mux, h.AddItem)
	MountUpdateItemHandler(mux, h.UpdateItem)
	MountRemoveItemHandler(mux, h.RemoveItem)
}

// Mount configures the mux to serve the front endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountShowCharacterHandler configures the mux to serve the "front" service
// "show_character" endpoint.
func MountShowCharacterHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/character/{id}", f)
}

// NewShowCharacterHandler creates a HTTP handler which loads the HTTP request
// and calls the "front" service "show_character" endpoint.
func NewShowCharacterHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeShowCharacterRequest(mux, decoder)
		encodeResponse = EncodeShowCharacterResponse(encoder)
		encodeError    = EncodeShowCharacterError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show_character")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListCharactersHandler configures the mux to serve the "front" service
// "list_characters" endpoint.
func MountListCharactersHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/character", f)
}

// NewListCharactersHandler creates a HTTP handler which loads the HTTP request
// and calls the "front" service "list_characters" endpoint.
func NewListCharactersHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeListCharactersResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list_characters")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountAddCharacterHandler configures the mux to serve the "front" service
// "add_character" endpoint.
func MountAddCharacterHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/character", f)
}

// NewAddCharacterHandler creates a HTTP handler which loads the HTTP request
// and calls the "front" service "add_character" endpoint.
func NewAddCharacterHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAddCharacterRequest(mux, decoder)
		encodeResponse = EncodeAddCharacterResponse(encoder)
		encodeError    = EncodeAddCharacterError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add_character")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateCharacterHandler configures the mux to serve the "front" service
// "update_character" endpoint.
func MountUpdateCharacterHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/character/{id}", f)
}

// NewUpdateCharacterHandler creates a HTTP handler which loads the HTTP
// request and calls the "front" service "update_character" endpoint.
func NewUpdateCharacterHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateCharacterRequest(mux, decoder)
		encodeResponse = EncodeUpdateCharacterResponse(encoder)
		encodeError    = EncodeUpdateCharacterError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update_character")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountRemoveCharacterHandler configures the mux to serve the "front" service
// "remove_character" endpoint.
func MountRemoveCharacterHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/character/{id}", f)
}

// NewRemoveCharacterHandler creates a HTTP handler which loads the HTTP
// request and calls the "front" service "remove_character" endpoint.
func NewRemoveCharacterHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRemoveCharacterRequest(mux, decoder)
		encodeResponse = EncodeRemoveCharacterResponse(encoder)
		encodeError    = EncodeRemoveCharacterError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "remove_character")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountShowInventoryHandler configures the mux to serve the "front" service
// "show_inventory" endpoint.
func MountShowInventoryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/inventory/{id}", f)
}

// NewShowInventoryHandler creates a HTTP handler which loads the HTTP request
// and calls the "front" service "show_inventory" endpoint.
func NewShowInventoryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeShowInventoryRequest(mux, decoder)
		encodeResponse = EncodeShowInventoryResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show_inventory")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateInventoryHandler configures the mux to serve the "front" service
// "update_inventory" endpoint.
func MountUpdateInventoryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/inventory/{id}/item/{item_id}", f)
}

// NewUpdateInventoryHandler creates a HTTP handler which loads the HTTP
// request and calls the "front" service "update_inventory" endpoint.
func NewUpdateInventoryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateInventoryRequest(mux, decoder)
		encodeResponse = EncodeUpdateInventoryResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update_inventory")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountRemoveInventoryHandler configures the mux to serve the "front" service
// "remove_inventory" endpoint.
func MountRemoveInventoryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/inventory/{id}/item/{item_id}", f)
}

// NewRemoveInventoryHandler creates a HTTP handler which loads the HTTP
// request and calls the "front" service "remove_inventory" endpoint.
func NewRemoveInventoryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRemoveInventoryRequest(mux, decoder)
		encodeResponse = EncodeRemoveInventoryResponse(encoder)
		encodeError    = EncodeRemoveInventoryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "remove_inventory")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountShowItemHandler configures the mux to serve the "front" service
// "show_item" endpoint.
func MountShowItemHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/item/{id}", f)
}

// NewShowItemHandler creates a HTTP handler which loads the HTTP request and
// calls the "front" service "show_item" endpoint.
func NewShowItemHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeShowItemRequest(mux, decoder)
		encodeResponse = EncodeShowItemResponse(encoder)
		encodeError    = EncodeShowItemError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show_item")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListItemsHandler configures the mux to serve the "front" service
// "list_items" endpoint.
func MountListItemsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/item", f)
}

// NewListItemsHandler creates a HTTP handler which loads the HTTP request and
// calls the "front" service "list_items" endpoint.
func NewListItemsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeListItemsResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list_items")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountAddItemHandler configures the mux to serve the "front" service
// "add_item" endpoint.
func MountAddItemHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/item", f)
}

// NewAddItemHandler creates a HTTP handler which loads the HTTP request and
// calls the "front" service "add_item" endpoint.
func NewAddItemHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAddItemRequest(mux, decoder)
		encodeResponse = EncodeAddItemResponse(encoder)
		encodeError    = EncodeAddItemError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add_item")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateItemHandler configures the mux to serve the "front" service
// "update_item" endpoint.
func MountUpdateItemHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/item/{id}", f)
}

// NewUpdateItemHandler creates a HTTP handler which loads the HTTP request and
// calls the "front" service "update_item" endpoint.
func NewUpdateItemHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateItemRequest(mux, decoder)
		encodeResponse = EncodeUpdateItemResponse(encoder)
		encodeError    = EncodeUpdateItemError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update_item")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountRemoveItemHandler configures the mux to serve the "front" service
// "remove_item" endpoint.
func MountRemoveItemHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/item/{id}", f)
}

// NewRemoveItemHandler creates a HTTP handler which loads the HTTP request and
// calls the "front" service "remove_item" endpoint.
func NewRemoveItemHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRemoveItemRequest(mux, decoder)
		encodeResponse = EncodeRemoveItemResponse(encoder)
		encodeError    = EncodeRemoveItemError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "remove_item")
		ctx = context.WithValue(ctx, goa.ServiceKey, "front")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
