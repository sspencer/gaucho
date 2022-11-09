# Goa Game API Example

This code is a test to design an API with a REST front end and backend GRPC microservices.

To run:

    go run ./services/front/cmd/front_server

While the micro services can run indepently, the front server currently embeds them.

Addresses:

* frontend: http://localhost:8000
* character: grpc://localhost:8080
* item: grpc://localhost:8081
* inventory: grpc://localhost:8082

Front API:

* ShowCharacter mounted on GET /character/{id}
* ListCharacters mounted on GET /character
* AddCharacter mounted on POST /character
* UpdateCharacter mounted on PUT /character/{id}
* RemoveCharacter mounted on DELETE /character/{id}
* ShowInventory mounted on GET /inventory/{id}
* UpdateInventory mounted on PUT /inventory/{id}/item/{item_id}
* RemoveInventory mounted on DELETE /inventory/{id}/item/{item_id}
* ShowItem mounted on GET /item/{id}
* ListItems mounted on GET /item
* AddItem mounted on POST /item
* UpdateItem mounted on PUT /item/{id}
* RemoveItem mounted on DELETE /item/{id}

Example requests in HTTP Client format are in the [requests.http](requests.http) file.