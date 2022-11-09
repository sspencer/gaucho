package design

import (
	. "gaucho/design"
	. "goa.design/goa/v3/dsl"
)

var _ = API("front", func() {
	Title("Gaucho Game Service API")
	Description("The gaucho game service API tracks gaming entities including characters, items and inventories.")
	Server("FrontServer", func() {
		Services("front")
		Host("development", func() {
			Description("Development hosts.")
			URI("http://localhost:8000/")
		})
	})
})

var _ = Service("front", func() {
	Description("Public HTTP frontend")

	Method("show_character", func() {
		Description("Get character by ID")
		Payload(func() {
			Field(1, "id", UInt32, "ID of character to retrieve", func() {
				Minimum(1)
			})
			Required("id")
		})
		Result(StoredCharacter)
		Error("not_found", ErrorResult, "Character not found")
		HTTP(func() {
			GET("/character/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("list_characters", func() {
		Description("Get all characters")
		Result(ArrayOf(StoredCharacter))
		HTTP(func() {
			GET("/character")
			Response(StatusOK)
		})
	})

	Method("add_character", func() {
		Description("Create new character")
		Payload(Character)
		Result(StoredCharacter)
		Error("already_exists", ErrorResult, "That name already exists")
		HTTP(func() {
			POST("/character")
			Response(StatusOK)
			Response("already_exists", StatusBadRequest)
		})
	})

	Method("update_character", func() {
		Description("Update new character")
		Payload(func() {
			Field(1, "id", UInt32, "ID of character to update", func() {
				Minimum(1)
			})
			Field(2, "character", Character, "Character values to update")
			Required("id", "character")
		})
		Result(StoredCharacter)
		Error("not_found", ErrorResult, "Character not found")
		Error("already_exists", ErrorResult, "That name already exists")
		HTTP(func() {
			PUT("/character/{id}")
			Body("character")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("already_exists", StatusBadRequest)
		})
	})

	Method("remove_character", func() {
		Description("Remove character from storage")
		Payload(func() {
			Field(1, "id", UInt32, "ID of character to remove", func() {
				Minimum(1)
			})
			Required("id")
		})
		Error("not_found", ErrorResult, "Character not found")
		HTTP(func() {
			DELETE("/character/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})

	//----------
	// inventory
	//----------
	Method("show_inventory", func() {
		Description("Show all items held by a character")
		Payload(func() {
			Field(1, "id", UInt32, "Character ID", func() {
				Minimum(1)
			})
			Required("id")
		})
		Result(ArrayOf(UInt32))
		HTTP(func() {
			GET("/inventory/{id}")
			Response(StatusOK)
		})
	})

	Method("update_inventory", func() {
		Description("Add item to character's inventory")
		Payload(func() {
			Field(1, "id", UInt32, "Character ID", func() {
				Minimum(1)
			})
			Field(2, "item_id", UInt32, "Item ID", func() {
				Minimum(1)
			})
			Required("id", "item_id")
		})
		Result(ArrayOf(UInt32))
		HTTP(func() {
			PUT("/inventory/{id}/item/{item_id}")
			Response(StatusOK)
		})
	})

	Method("remove_inventory", func() {
		Description("Remove item from character's inventory")
		Payload(func() {
			Field(1, "id", UInt32, "Character ID", func() {
				Minimum(1)
			})
			Field(2, "item_id", UInt32, "Item ID", func() {
				Minimum(1)
			})
			Required("id", "item_id")
		})
		Error("not_found", ErrorResult, "Character or Item not found")
		HTTP(func() {
			DELETE("/inventory/{id}/item/{item_id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})

	//----------
	// items
	//----------
	Method("show_item", func() {
		Description("Get item by ID")
		Payload(func() {
			Field(1, "id", UInt32, "ID of item to retrieve", func() {
				Minimum(1)
			})
			Required("id")
		})
		Result(StoredItem)
		Error("not_found", ErrorResult, "Item not found")
		HTTP(func() {
			GET("/item/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("list_items", func() {
		Description("Get all items")
		Result(ArrayOf(StoredItem))
		HTTP(func() {
			GET("/item")
			Response(StatusOK)
		})
	})

	Method("add_item", func() {
		Description("Create new item")
		Payload(Item)
		Result(StoredItem)
		Error("already_exists", ErrorResult, "That name already exists")
		HTTP(func() {
			POST("/item")
			Response(StatusOK)
			Response("already_exists", StatusBadRequest)
		})
	})

	Method("update_item", func() {
		Description("Update new item")
		Payload(func() {
			Field(1, "id", UInt32, "ID of item to update", func() {
				Minimum(1)
			})
			Field(2, "item", Item, "Item values to update")
			Required("id", "item")
		})
		Result(StoredItem)
		Error("not_found", ErrorResult, "Item not found")
		Error("already_exists", ErrorResult, "That name already exists")
		HTTP(func() {
			PUT("/item/{id}")
			Body("item")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("already_exists", StatusBadRequest)
		})
	})

	Method("remove_item", func() {
		Description("Remove item from storage")
		Payload(func() {
			Field(1, "id", UInt32, "ID of item to remove", func() {
				Minimum(1)
			})
			Required("id")
		})
		Error("not_found", ErrorResult, "Item not found")
		HTTP(func() {
			DELETE("/item/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})

})
