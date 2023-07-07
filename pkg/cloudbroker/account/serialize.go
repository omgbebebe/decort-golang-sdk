package account

import (
	"encoding/json"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/serialization"
)

// Serialize returns JSON-serialized []byte. Used as a wrapper over json.Marshal and json.MarshalIndent functions.
//
// In order to serialize with indent make sure to follow these guidelines:
//   - First argument -> prefix
//   - Second argument -> indent
func (la ListAccounts) Serialize(params ...string) (serialization.Serialized, error) {
	if la.EntryCount == 0 {
		return []byte{}, nil
	}

	if len(params) > 1 {
		prefix := params[0]
		indent := params[1]

		return json.MarshalIndent(la, prefix, indent)
	}

	return json.Marshal(la)
}

// Serialize returns JSON-serialized []byte. Used as a wrapper over json.Marshal and json.MarshalIndent functions.
//
// In order to serialize with indent make sure to follow these guidelines:
//   - First argument -> prefix
//   - Second argument -> indent
func (ia ItemAccount) Serialize(params ...string) (serialization.Serialized, error) {
	if len(params) > 1 {
		prefix := params[0]
		indent := params[1]

		return json.MarshalIndent(ia, prefix, indent)
	}

	return json.Marshal(ia)
}
