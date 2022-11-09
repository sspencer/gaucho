package frontapi

import (
	"gaucho/services/character/gen/character"
	"gaucho/services/front/gen/front"
	"gaucho/services/item/gen/item"
)

func toFrontStoredCharacter(c *character.StoredCharacter) *front.StoredCharacter {
	return &front.StoredCharacter{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Experience:  c.Experience,
		Health:      c.Health,
	}
}

func toFrontStoredCharacterList(s []*character.StoredCharacter) []*front.StoredCharacter {
	f := make([]*front.StoredCharacter, len(s))
	for i := range f {
		f[i] = toFrontStoredCharacter(s[i])
	}

	return f
}

func toCharacter(c *front.Character) *character.Character {
	return &character.Character{
		Name:        c.Name,
		Description: c.Description,
		Experience:  c.Experience,
		Health:      c.Health,
	}
}

func toFrontStoredItem(i *item.StoredItem) *front.StoredItem {
	return &front.StoredItem{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Protection:  i.Protection,
		Healing:     i.Healing,
		Damage:      i.Damage,
	}
}

func toFrontStoredItemList(s []*item.StoredItem) []*front.StoredItem {
	f := make([]*front.StoredItem, len(s))
	for i := range f {
		f[i] = toFrontStoredItem(s[i])
	}

	return f
}

func toItem(i *front.Item) *item.Item {
	return &item.Item{
		Name:        i.Name,
		Description: i.Description,
		Protection:  i.Protection,
		Healing:     i.Healing,
		Damage:      i.Damage,
	}
}