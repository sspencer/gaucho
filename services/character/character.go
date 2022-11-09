package characterapi

import (
	"context"
	"fmt"
	character "gaucho/services/character/gen/character"
	"golang.org/x/exp/maps"
	"log"
	"sync/atomic"
)

// character service example implementation.
// The example methods log the requests and return zero values.
type charactersrvc struct {
	logger *log.Logger
	store  map[uint32]*character.StoredCharacter
	nextID uint32
}

// NewCharacter returns the character service implementation.
func NewCharacter(logger *log.Logger) character.Service {
	return &charactersrvc{
		logger: logger,
		store:  make(map[uint32]*character.StoredCharacter),
		nextID: 0,
	}
}

// Retrieve character by the given id
func (s *charactersrvc) Get(ctx context.Context, p uint32) (*character.StoredCharacter, error) {
	s.logger.Print("character.get")

	res, ok := s.store[p]
	if !ok {
		return nil, character.NotFound(fmt.Sprintf("Character %d not found", p))
	}

	return res, nil
}

// Retrieve all characters
func (s *charactersrvc) GetAll(ctx context.Context) ([]*character.StoredCharacter, error) {
	s.logger.Print("character.get_all")
	return maps.Values(s.store), nil
}

// Create a character
func (s *charactersrvc) Add(ctx context.Context, p *character.Character) (*character.StoredCharacter, error) {
	s.logger.Print("character.add")

	if s.nameExists(p.Name, 0) {
		return nil, character.AlreadyExists(fmt.Sprintf("Character %q already exists", p.Name))
	}

	atomic.AddUint32(&s.nextID, 1)

	res := &character.StoredCharacter{
		ID:          s.nextID,
		Name:        p.Name,
		Description: p.Description,
		Health:      p.Health,
		Experience:  p.Experience,
	}

	s.store[s.nextID] = res

	return res, nil
}

// Update a character
func (s *charactersrvc) Update(ctx context.Context, p *character.UpdatePayload) (*character.StoredCharacter, error) {
	s.logger.Print("character.update")

	res, ok := s.store[p.ID]
	if !ok {
		return nil, character.NotFound(fmt.Sprintf("Character %d not found", p.ID))
	}

	if s.nameExists(p.Character.Name, p.ID) {
		return nil, character.AlreadyExists(fmt.Sprintf("Character %q already exists", p.Character.Name))
	}

	res.Name = p.Character.Name
	res.Description = p.Character.Description
	res.Health = p.Character.Health
	res.Experience = p.Character.Experience

	s.store[res.ID] = res

	return res, nil
}

// Remove character
func (s *charactersrvc) Remove(ctx context.Context, p *character.RemovePayload) (err error) {
	s.logger.Print("character.remove")

	res, ok := s.store[p.ID]
	if !ok {
		return character.NotFound(fmt.Sprintf("Character %d not found", p.ID))
	}

	delete(s.store, res.ID)

	return nil
}

func (s *charactersrvc) nameExists(name string, id uint32) bool {
	for _, c := range s.store {
		// check for name clashes when adding (id == 0),
		// skipping when object matches itself
		if (id == 0 || id != c.ID) && name == c.Name {
			return true
		}
	}

	return false
}
