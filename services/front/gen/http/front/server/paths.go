// Code generated by goa v3.10.2, DO NOT EDIT.
//
// HTTP request path constructors for the front service.
//
// Command:
// $ goa gen gaucho/services/front/design -o services/front

package server

import (
	"fmt"
)

// ShowCharacterFrontPath returns the URL path to the front service show_character HTTP endpoint.
func ShowCharacterFrontPath(id uint32) string {
	return fmt.Sprintf("/character/%v", id)
}

// ListCharactersFrontPath returns the URL path to the front service list_characters HTTP endpoint.
func ListCharactersFrontPath() string {
	return "/character"
}

// AddCharacterFrontPath returns the URL path to the front service add_character HTTP endpoint.
func AddCharacterFrontPath() string {
	return "/character"
}

// UpdateCharacterFrontPath returns the URL path to the front service update_character HTTP endpoint.
func UpdateCharacterFrontPath(id uint32) string {
	return fmt.Sprintf("/character/%v", id)
}

// RemoveCharacterFrontPath returns the URL path to the front service remove_character HTTP endpoint.
func RemoveCharacterFrontPath(id uint32) string {
	return fmt.Sprintf("/character/%v", id)
}

// ShowInventoryFrontPath returns the URL path to the front service show_inventory HTTP endpoint.
func ShowInventoryFrontPath(id uint32) string {
	return fmt.Sprintf("/inventory/%v", id)
}

// UpdateInventoryFrontPath returns the URL path to the front service update_inventory HTTP endpoint.
func UpdateInventoryFrontPath(id uint32, itemID uint32) string {
	return fmt.Sprintf("/inventory/%v/item/%v", id, itemID)
}

// RemoveInventoryFrontPath returns the URL path to the front service remove_inventory HTTP endpoint.
func RemoveInventoryFrontPath(id uint32, itemID uint32) string {
	return fmt.Sprintf("/inventory/%v/item/%v", id, itemID)
}

// ShowItemFrontPath returns the URL path to the front service show_item HTTP endpoint.
func ShowItemFrontPath(id uint32) string {
	return fmt.Sprintf("/item/%v", id)
}

// ListItemsFrontPath returns the URL path to the front service list_items HTTP endpoint.
func ListItemsFrontPath() string {
	return "/item"
}

// AddItemFrontPath returns the URL path to the front service add_item HTTP endpoint.
func AddItemFrontPath() string {
	return "/item"
}

// UpdateItemFrontPath returns the URL path to the front service update_item HTTP endpoint.
func UpdateItemFrontPath(id uint32) string {
	return fmt.Sprintf("/item/%v", id)
}

// RemoveItemFrontPath returns the URL path to the front service remove_item HTTP endpoint.
func RemoveItemFrontPath(id uint32) string {
	return fmt.Sprintf("/item/%v", id)
}