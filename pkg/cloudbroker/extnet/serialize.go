package extnet

import (
	"encoding/json"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/serialization"
)

// Serialize returns JSON-serialized []byte. Used as a wrapper over json.Marshal and json.MarshalIndent functions.
//
// In order to serialize with indent make sure to follow these guidelines:
//   - First argument -> prefix
//   - Second argument -> indent
func (lenet ListExtNet) Serialize(params ...string) (serialization.Serialized, error) {
	if len(lenet.Data) == 0 {
		return []byte{}, nil
	}

	if len(params) > 1 {
		prefix := params[0]
		indent := params[1]

		return json.MarshalIndent(lenet, prefix, indent)
	}

	return json.Marshal(lenet)
}

// Serialize returns JSON-serialized []byte. Used as a wrapper over json.Marshal and json.MarshalIndent functions.
//
// In order to serialize with indent make sure to follow these guidelines:
//   - First argument -> prefix
//   - Second argument -> indent
func (ienet ItemExtNet) Serialize(params ...string) (serialization.Serialized, error) {
	if len(params) > 1 {
		prefix := params[0]
		indent := params[1]

		return json.MarshalIndent(ienet, prefix, indent)
	}

	return json.Marshal(ienet)
}
