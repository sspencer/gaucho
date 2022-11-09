// Code generated by goa v3.10.2, DO NOT EDIT.
//
// character gRPC client CLI support package
//
// Command:
// $ goa gen gaucho/services/character/design -o services/character

package client

import (
	"encoding/json"
	"fmt"
	character "gaucho/services/character/gen/character"
	characterpb "gaucho/services/character/gen/grpc/character/pb"
)

// BuildGetPayload builds the payload for the character get endpoint from CLI
// flags.
func BuildGetPayload(characterGetMessage string) (uint32, error) {
	var err error
	var message characterpb.GetRequest
	{
		if characterGetMessage != "" {
			err = json.Unmarshal([]byte(characterGetMessage), &message)
			if err != nil {
				var zero uint32
				return zero, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"field\": 2931724321\n   }'")
			}
		}
	}
	v := message.Field
	return v, nil
}

// BuildAddPayload builds the payload for the character add endpoint from CLI
// flags.
func BuildAddPayload(characterAddMessage string) (*character.Character, error) {
	var err error
	var message characterpb.AddRequest
	{
		if characterAddMessage != "" {
			err = json.Unmarshal([]byte(characterAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Quo odit rerum et.\",\n      \"experience\": 56,\n      \"health\": 12,\n      \"name\": \"tvt\"\n   }'")
			}
		}
	}
	v := &character.Character{
		Name:        message.Name,
		Description: message.Description,
		Health:      int(message.Health),
		Experience:  int(message.Experience),
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the character update endpoint from
// CLI flags.
func BuildUpdatePayload(characterUpdateMessage string) (*character.UpdatePayload, error) {
	var err error
	var message characterpb.UpdateRequest
	{
		if characterUpdateMessage != "" {
			err = json.Unmarshal([]byte(characterUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"character\": {\n         \"description\": \"Ab enim dolorem quod iste.\",\n         \"experience\": 42,\n         \"health\": 99,\n         \"name\": \"r1t\"\n      },\n      \"id\": 3972646431\n   }'")
			}
		}
	}
	v := &character.UpdatePayload{
		ID: message.Id,
	}
	if message.Character != nil {
		v.Character = protobufCharacterpbCharacter2ToCharacterCharacter(message.Character)
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the character remove endpoint from
// CLI flags.
func BuildRemovePayload(characterRemoveMessage string) (*character.RemovePayload, error) {
	var err error
	var message characterpb.RemoveRequest
	{
		if characterRemoveMessage != "" {
			err = json.Unmarshal([]byte(characterRemoveMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 3796969755\n   }'")
			}
		}
	}
	v := &character.RemovePayload{
		ID: message.Id,
	}

	return v, nil
}
