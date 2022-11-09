package itemapi

import (
	"context"
	"fmt"
	item "gaucho/services/item/gen/item"
	"golang.org/x/exp/maps"
	"log"
	"sync/atomic"
)

// item service example implementation.
// The example methods log the requests and return zero values.
type itemsrvc struct {
	logger *log.Logger
	store  map[uint32]*item.StoredItem
	nextID uint32
}

// NewItem returns the item service implementation.
func NewItem(logger *log.Logger) item.Service {
	return &itemsrvc{
		logger: logger,
		store:  make(map[uint32]*item.StoredItem),
		nextID: 0,
	}
}

// Retrieve item by the given id
func (s *itemsrvc) Get(ctx context.Context, p uint32) (*item.StoredItem, error) {
	s.logger.Print("item.get")
	res, ok := s.store[p]
	if !ok {
		return nil, item.NotFound(fmt.Sprintf("Item %d not found", p))
	}

	return res, nil
}

// Retrieve all items
func (s *itemsrvc) GetAll(ctx context.Context) ([]*item.StoredItem, error) {
	s.logger.Print("item.get_all")
	return maps.Values(s.store), nil
}

// Create an item
func (s *itemsrvc) Add(ctx context.Context, p *item.Item) (res *item.StoredItem, err error) {
	s.logger.Print("item.add")
	if s.nameExists(p.Name, 0) {
		return nil, item.AlreadyExists(fmt.Sprintf("Item %q already exists", p.Name))
	}

	atomic.AddUint32(&s.nextID, 1)

	res = &item.StoredItem{
		ID:          s.nextID,
		Name:        p.Name,
		Description: p.Description,
		Damage:      p.Damage,
		Healing:     p.Healing,
		Protection:  p.Protection,
	}

	s.store[s.nextID] = res

	return
}

// Update an item
func (s *itemsrvc) Update(ctx context.Context, p *item.UpdatePayload) (*item.StoredItem, error) {
	s.logger.Print("item.update")
	res, ok := s.store[p.ID]
	if !ok {
		return nil, item.NotFound(fmt.Sprintf("Item %d not found", p.ID))
	}

	if s.nameExists(p.Item.Name, p.ID) {
		return nil, item.AlreadyExists(fmt.Sprintf("Item %q already exists", p.Item.Name))
	}
	res.Name = p.Item.Name
	res.Description = p.Item.Description
	res.Healing = p.Item.Healing
	res.Damage = p.Item.Damage
	res.Protection = p.Item.Protection

	s.store[res.ID] = res

	return res, nil
}

// Remove item
func (s *itemsrvc) Remove(ctx context.Context, p *item.RemovePayload) (err error) {
	s.logger.Print("item.remove")
	res, ok := s.store[p.ID]
	if !ok {
		return item.NotFound(fmt.Sprintf("Item %d not found", p.ID))
	}

	delete(s.store, res.ID)

	return nil
}

func (s *itemsrvc) nameExists(name string, id uint32) bool {
	for _, c := range s.store {
		// check for name clashes when adding (id == 0),
		// skipping when object matches itself
		if (id == 0 || id != c.ID) && name == c.Name {
			return true
		}
	}

	return false
}
