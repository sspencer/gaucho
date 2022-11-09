package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("inventory", func() {
	Title("Gaucho Inventory Service API")
	Description("The gaucho inventory service API provides methods to store items with characters.")
	Server("InventoryServer", func() {
		Services("inventory")
		Host("development", func() {
			Description("Development hosts.")
			URI("grpc://localhost:8082/")
		})
	})
})

var _ = Service("inventory", func() {
	Description("Service that provides character API")

	Method("get", func() {
		Description("Retrieve all inventory items for the given character")
		Payload(UInt32, "character_id", func() {
			Minimum(1)
		})
		Result(ArrayOf(UInt32))
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("add", func() {
		Description("Add inventory item to the given character")
		Payload(func() {
			Field(1, "character_id", UInt32, "Character ID", func() {
				Minimum(1)
			})
			Field(2, "item_id", UInt32, "Item ID to add to character's inventory", func() {
				Minimum(1)
			})
			Required("character_id", "item_id")
		})
		Result(ArrayOf(UInt32))
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove inventory item from the given character")
		Payload(func() {
			Field(1, "character_id", UInt32, "Character ID", func() {
				Minimum(1)
			})
			Field(2, "item_id", UInt32, "Item ID to add to character's inventory", func() {
				Minimum(1)
			})
			Required("character_id", "item_id")
		})
		Error("not_found", String, "Entity not found")
		GRPC(func() {
			Response(CodeOK)
			Response(CodeNotFound)
		})
	})

})
