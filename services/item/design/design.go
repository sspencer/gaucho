package design

import (
	. "gaucho/design"
	. "goa.design/goa/v3/dsl"
)

var _ = API("item", func() {
	Title("Gaucho Item Service API")
	Description("The gaucho item service API provides methods to store and manipulate items.")
	Server("ItemServer", func() {
		Services("item")
		Host("development", func() {
			Description("Development hosts.")
			URI("grpc://localhost:8081/")
		})
	})
})

var _ = Service("item", func() {
	Description("Service that provides item API")

	Method("get", func() {
		Description("Retrieve item by the given id")
		Payload(UInt32, "id", func() {
			Minimum(1)
		})
		Result(StoredItem)
		Error("not_found", String, "Item not found")
		GRPC(func() {
			Response(CodeOK)
			Response(CodeNotFound)
		})
	})

	Method("get_all", func() {
		Description("Retrieve all items")
		Result(ArrayOf(StoredItem))
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("add", func() {
		Description("Create an item")
		Payload(Item)
		Error("already_exists", String, "Item name already exists")
		Result(StoredItem)
		GRPC(func() {
			Response(CodeOK)
			Response("already_exists", CodeAlreadyExists)
		})
	})

	Method("update", func() {
		Description("Update an item")
		Payload(func() {
			Field(1, "id", UInt32, "ID of item to update", func() {
				Minimum(1)
			})
			Field(2, "item", Item, "Item fields to update")
			Required("id", "item")
		})
		Result(StoredItem)
		Error("already_exists", String, "Item name already exists")
		Error("not_found", String, "Item not found")
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
			Response("already_exists", CodeAlreadyExists)
		})
	})

	Method("remove", func() {
		Description("Remove item")
		Payload(func() {
			Field(1, "id", UInt32, "ID of item to remove", func() {
				Minimum(1)
			})
			Required("id")
		})
		Error("not_found", String, "Item not found")
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})
})
