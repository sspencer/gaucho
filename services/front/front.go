package frontapi

import (
	"context"
	"gaucho/services/character/gen/character"
	front "gaucho/services/front/gen/front"
	"gaucho/services/inventory/gen/inventory"
	"gaucho/services/item/gen/item"
	"log"
	"strings"
)

// front service example implementation.
// The example methods log the requests and return zero values.
type frontsrvc struct {
	logger    *log.Logger
	character character.Service
	item      item.Service
	inventory inventory.Service
}

// NewFront returns the front service implementation.
func NewFront(logger *log.Logger, charSvc character.Service, itemSvc item.Service, inventorySvc inventory.Service) front.Service {
	return &frontsrvc{
		logger:    logger,
		character: charSvc,
		item:      itemSvc,
		inventory: inventorySvc,
	}
}

// Get character by ID
func (s *frontsrvc) ShowCharacter(ctx context.Context, p *front.ShowCharacterPayload) (*front.StoredCharacter, error) {
	s.logger.Print("front.show_character")
	res, err := s.character.Get(ctx, p.ID)
	if err != nil {
		return nil, front.MakeNotFound(err)
	}

	return toFrontStoredCharacter(res), nil
}

// Get all characters
func (s *frontsrvc) ListCharacters(ctx context.Context) ([]*front.StoredCharacter, error) {
	s.logger.Print("front.list_characters")
	res, _ := s.character.GetAll(ctx)
	return toFrontStoredCharacterList(res), nil
}

// Create new character
func (s *frontsrvc) AddCharacter(ctx context.Context, p *front.Character) (*front.StoredCharacter, error) {
	s.logger.Print("front.add_character")
	res, err := s.character.Add(ctx, toCharacter(p))
	if err != nil {
		return nil, front.MakeAlreadyExists(err)
	}
	return toFrontStoredCharacter(res), nil
}

// Update new character
func (s *frontsrvc) UpdateCharacter(ctx context.Context, p *front.UpdateCharacterPayload) (*front.StoredCharacter, error) {
	s.logger.Print("front.update_character")
	payload := &character.UpdatePayload{
		ID:        p.ID,
		Character: toCharacter(p.Character),
	}
	res, err := s.character.Update(ctx, payload)
	if err != nil {
		// not the best approach, but it will have to do for now
		if strings.Contains(err.Error(), "already") {
			return nil, front.MakeAlreadyExists(err)
		}

		return nil, front.MakeNotFound(err)
	}

	return toFrontStoredCharacter(res), nil
}

// Remove character from storage
func (s *frontsrvc) RemoveCharacter(ctx context.Context, p *front.RemoveCharacterPayload) error {
	s.logger.Print("front.remove_character")
	payload := &character.RemovePayload{ID: p.ID}
	err := s.character.Remove(ctx, payload)
	if err != nil {
		return front.MakeNotFound(err)
	}
	return nil
}

// Get inventory of given character
func (s *frontsrvc) ShowInventory(ctx context.Context, p *front.ShowInventoryPayload) ([]uint32, error) {
	s.logger.Print("front.show_inventory")

	// verify character exists before getting inventory
	_, err := s.character.Get(ctx, p.ID)
	if err != nil {
		return nil, front.MakeNotFound(err)
	}

	res, _ := s.inventory.Get(ctx, p.ID)
	return res, nil
}

// Add item to inventory of given character
func (s *frontsrvc) UpdateInventory(ctx context.Context, p *front.UpdateInventoryPayload) ([]uint32, error) {
	s.logger.Print("front.update_inventory")

	// verify character exists before adding to inventory
	_, err := s.character.Get(ctx, p.ID)
	if err != nil {
		return nil, front.MakeNotFound(err)
	}

	payload := &inventory.AddPayload{
		CharacterID: p.ID,
		ItemID:      p.ItemID,
	}
	res, _ := s.inventory.Add(ctx, payload)
	return res, nil
}

// Remove item from inventory of given character
func (s *frontsrvc) RemoveInventory(ctx context.Context, p *front.RemoveInventoryPayload) error {
	s.logger.Print("front.remove_inventory")

	// verify character exists trying to delete from inventory
	_, err := s.character.Get(ctx, p.ID)
	if err != nil {
		return front.MakeNotFound(err)
	}

	payload := &inventory.RemovePayload{
		CharacterID: p.ID,
		ItemID:      p.ItemID,
	}

	err = s.inventory.Remove(ctx, payload)
	if err != nil {
		return front.MakeNotFound(err)
	}
	return nil
}

// Get item by ID
func (s *frontsrvc) ShowItem(ctx context.Context, p *front.ShowItemPayload) (*front.StoredItem, error) {
	s.logger.Print("front.show_item")
	res, err := s.item.Get(ctx, p.ID)
	if err != nil {
		return nil, front.MakeNotFound(err)
	}

	return toFrontStoredItem(res), nil
}

// Get all items
func (s *frontsrvc) ListItems(ctx context.Context) ([]*front.StoredItem, error) {
	s.logger.Print("front.list_items")
	res, _ := s.item.GetAll(ctx)
	return toFrontStoredItemList(res), nil
}

// Create new item
func (s *frontsrvc) AddItem(ctx context.Context, p *front.Item) (*front.StoredItem, error) {
	s.logger.Print("front.add_item")
	res, err := s.item.Add(ctx, toItem(p))
	if err != nil {
		return nil, front.MakeAlreadyExists(err)
	}

	return toFrontStoredItem(res), nil
}

// Update new item
func (s *frontsrvc) UpdateItem(ctx context.Context, p *front.UpdateItemPayload) (*front.StoredItem, error) {
	s.logger.Print("front.update_item")
	payload := &item.UpdatePayload{
		ID:   p.ID,
		Item: toItem(p.Item),
	}

	res, err := s.item.Update(ctx, payload)
	if err != nil {
		// not the best approach, but it will have to do for now
		if strings.Contains(err.Error(), "already") {
			return nil, front.MakeAlreadyExists(err)
		}

		return nil, front.MakeNotFound(err)
	}

	return toFrontStoredItem(res), nil
}

// Remove item from storage
func (s *frontsrvc) RemoveItem(ctx context.Context, p *front.RemoveItemPayload) error {
	s.logger.Print("front.remove_item")
	payload := &item.RemovePayload{ID: p.ID}
	err := s.item.Remove(ctx, payload)
	if err != nil {
		return front.MakeNotFound(err)
	}
	return nil
}
