// Code generated by goa v3.10.2, DO NOT EDIT.
//
// front views
//
// Command:
// $ goa gen gaucho/services/front/design -o services/front

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// StoredCharacter is the viewed result type that is projected based on a view.
type StoredCharacter struct {
	// Type to project
	Projected *StoredCharacterView
	// View to render
	View string
}

// StoredItem is the viewed result type that is projected based on a view.
type StoredItem struct {
	// Type to project
	Projected *StoredItemView
	// View to render
	View string
}

// StoredCharacterView is a type that runs validations on a projected type.
type StoredCharacterView struct {
	// ID is the unique id of the character.
	ID *uint32
	// Unique Character Name
	Name *string
	// Character Description
	Description *string
	// Character Experience
	Experience *int
	// Character Health
	Health *int
}

// StoredItemView is a type that runs validations on a projected type.
type StoredItemView struct {
	// ID is the unique id of the item.
	ID *uint32
	// Unique Item Name
	Name *string
	// Item Description
	Description *string
	// Amount of damage item causes
	Damage *int
	// Amount of healing item generates
	Healing *int
	// Amount of protection item provides
	Protection *int
}

var (
	// StoredCharacterMap is a map indexing the attribute names of StoredCharacter
	// by view name.
	StoredCharacterMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"health",
			"experience",
		},
	}
	// StoredItemMap is a map indexing the attribute names of StoredItem by view
	// name.
	StoredItemMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"damage",
			"healing",
			"protection",
		},
	}
)

// ValidateStoredCharacter runs the validations defined on the viewed result
// type StoredCharacter.
func ValidateStoredCharacter(result *StoredCharacter) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredCharacterView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStoredItem runs the validations defined on the viewed result type
// StoredItem.
func ValidateStoredItem(result *StoredItem) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredItemView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStoredCharacterView runs the validations defined on
// StoredCharacterView using the "default" view.
func ValidateStoredCharacterView(result *StoredCharacterView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "result"))
	}
	if result.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 2, true))
		}
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 50, false))
		}
	}
	if result.Health != nil {
		if *result.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.health", *result.Health, 0, true))
		}
	}
	if result.Health != nil {
		if *result.Health > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.health", *result.Health, 100, false))
		}
	}
	if result.Experience != nil {
		if *result.Experience < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.experience", *result.Experience, 1, true))
		}
	}
	if result.Experience != nil {
		if *result.Experience > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.experience", *result.Experience, 100, false))
		}
	}
	return
}

// ValidateStoredItemView runs the validations defined on StoredItemView using
// the "default" view.
func ValidateStoredItemView(result *StoredItemView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "result"))
	}
	if result.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "result"))
	}
	if result.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 2, true))
		}
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 50, false))
		}
	}
	if result.Damage != nil {
		if *result.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.damage", *result.Damage, 0, true))
		}
	}
	if result.Damage != nil {
		if *result.Damage > 25 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.damage", *result.Damage, 25, false))
		}
	}
	if result.Healing != nil {
		if *result.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.healing", *result.Healing, 0, true))
		}
	}
	if result.Healing != nil {
		if *result.Healing > 50 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.healing", *result.Healing, 50, false))
		}
	}
	if result.Protection != nil {
		if *result.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.protection", *result.Protection, 0, true))
		}
	}
	if result.Protection != nil {
		if *result.Protection > 20 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.protection", *result.Protection, 20, false))
		}
	}
	return
}