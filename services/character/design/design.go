package design

import (
	. "gaucho/design"
	. "goa.design/goa/v3/dsl"
)

var _ = API("character", func() {
	Title("Gaucho Character Service API")
	Description("The gaucho character service API provides methods to store and manipulate characters.")
	Server("CharacterServer", func() {
		Services("character")
		Host("development", func() {
			Description("Development hosts.")
			URI("grpc://localhost:8080/")
		})
	})
})

var _ = Service("character", func() {
	Description("Service that provides character API")

	Method("get", func() {
		Description("Retrieve character by the given id")
		Payload(UInt32, "id", func() {
			Minimum(1)
		})
		Result(StoredCharacter)
		Error("not_found", String, "Character not found")
		GRPC(func() {
			Response(CodeOK)
			Response(CodeNotFound)
		})
	})

	Method("get_all", func() {
		Description("Retrieve all characters")
		Result(ArrayOf(StoredCharacter))
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("add", func() {
		Description("Create a character")
		Payload(Character)
		Error("already_exists", String, "Character name already exists")
		Result(StoredCharacter)
		GRPC(func() {
			Response(CodeOK)
			Response("already_exists", CodeAlreadyExists)
		})
	})

	Method("update", func() {
		Description("Update a character")
		Payload(func() {
			Field(1, "id", UInt32, "ID of character to update", func() {
				Minimum(1)
			})
			Field(2, "character", Character, "Character parameters to update")
			Required("id", "character")
		})
		Result(StoredCharacter)
		Error("already_exists", String, "Character name already exists")
		Error("not_found", String, "Character not found")
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
			Response("already_exists", CodeAlreadyExists)
		})
	})

	Method("remove", func() {
		Description("Remove character")
		Payload(func() {
			Field(1, "id", UInt32, "ID of character to remove", func() {
				Minimum(1)
			})
			Required("id")
		})
		Error("not_found", String, "Character not found")
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})
})
