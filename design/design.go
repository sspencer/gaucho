package design

import . "goa.design/goa/v3/dsl"

var Character = Type("Character", func() {
	Description("Character in the game")
	Field(1, "name", String, "Unique Character Name", func() {
		MinLength(2)
		MaxLength(50)
	})
	Field(2, "description", String, "Character Description")
	Field(3, "health", Int, "Character Health", func() {
		Minimum(0)
		Maximum(100)
	})
	Field(4, "experience", Int, "Character Experience", func() {
		Minimum(1)
		Maximum(100)
	})
	Required("name", "description", "health", "experience")
})

var StoredCharacter = ResultType("application/vnd.gaucho.stored-character", func() {
	Description("A StoredCharacter describes a character retrieved by the character service.")
	Reference(Character)
	ContentType("application/json") // Override Content-Type header
	TypeName("StoredCharacter")

	Attributes(func() {
		Attribute("id", UInt32, "ID is the unique id of the character.", func() {
			Example(42)
			Meta("rpc:tag", "1")
		})
		Field(2, "name")
		Field(3, "description")
		Field(4, "experience")
		Field(5, "health")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("health")
		Attribute("experience")
	})

	Required("id", "description", "name", "health", "experience")
})

var Item = Type("Item", func() {
	Description("A game item like a sword, shield or potion.")
	Field(1, "name", String, "Unique Item Name", func() {
		MinLength(2)
		MaxLength(50)
	})
	Field(2, "description", String, "Item Description")
	Field(3, "damage", Int, "Amount of damage item causes", func() {
		Minimum(0)
		Maximum(25)
	})
	Field(4, "healing", Int, "Amount of healing item generates", func() {
		Minimum(0)
		Maximum(50)
	})
	Field(5, "protection", Int, "Amount of protection item provides", func() {
		Minimum(0)
		Maximum(20)
	})
	Required("name", "damage", "healing", "protection")
})

var StoredItem = ResultType("application/vnd.gaucho.stored-item", func() {
	Description("A StoredItem describes a item retrieved by the character item.")
	Reference(Item)
	ContentType("application/json") // Override Content-Type header
	TypeName("StoredItem")

	Attributes(func() {
		Attribute("id", UInt32, "ID is the unique id of the item.", func() {
			Example(87)
			Meta("rpc:tag", "1")
		})
		Field(2, "name")
		Field(3, "description")
		Field(4, "damage")
		Field(5, "healing")
		Field(6, "protection")

	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("damage")
		Attribute("healing")
		Attribute("protection")
	})

	Required("id", "name", "damage", "healing", "protection")
})
