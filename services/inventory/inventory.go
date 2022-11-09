package inventoryapi

import (
	"context"
	inventory "gaucho/services/inventory/gen/inventory"
	"log"
)

// inventory service example implementation.
// The example methods log the requests and return zero values.
type inventorysrvc struct {
	logger *log.Logger
	store  map[uint32][]uint32
}

// NewInventory returns the inventory service implementation.
func NewInventory(logger *log.Logger) inventory.Service {
	return &inventorysrvc{
		logger: logger,
		store:  make(map[uint32][]uint32),
	}
}

// Retrieve all inventory items for the given character
func (s *inventorysrvc) Get(ctx context.Context, p uint32) ([]uint32, error) {
	s.logger.Print("inventory.get")

	res, _ := s.store[p]
	return res, nil
}

// Add inventory item to the given character
func (s *inventorysrvc) Add(ctx context.Context, p *inventory.AddPayload) ([]uint32, error) {
	s.logger.Print("inventory.add")
	res, _ := s.store[p.CharacterID]
	res = append(res, p.ItemID)
	s.store[p.CharacterID] = res
	return res, nil
}

// Remove inventory item from the given character
func (s *inventorysrvc) Remove(ctx context.Context, p *inventory.RemovePayload) error {
	s.logger.Print("inventory.remove")
	notFound := inventory.NotFound("Character or item doesn't exist in inventory")
	res, ok := s.store[p.CharacterID]
	if !ok {
		return notFound
	}

	index := indexOf(res, p.ItemID)
	if index == -1 {
		return notFound
	}

	s.store[p.CharacterID] = removeIndex(res, index)

	return nil
}

func indexOf(s []uint32, itemId uint32) int {
	for k, v := range s {
		if itemId == v {
			return k
		}
	}
	return -1
}

func removeIndex(s []uint32, index int) []uint32 {
	ret := make([]uint32, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
