package computeci

import (
	"encoding/json"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/serialization"
)

// Serialize returns JSON-serialized []byte. Used as a wrapper over json.Marshal and json.MarshalIndent functions.
//
// In order to serialize with indent make sure to follow these guidelines:
//   - First argument -> prefix
//   - Second argument -> indent
func (lci ListComputeCI) Serialize(params ...string) (serialization.Serialized, error) {
	if len(lci) == 0 {
		return []byte{}, nil
	}

	if len(params) > 1 {
		prefix := params[0]
		indent := params[1]

		return json.MarshalIndent(lci, prefix, indent)
	}

	return json.Marshal(lci)
}

// Serialize returns JSON-serialized []byte. Used as a wrapper over json.Marshal and json.MarshalIndent functions.
//
// In order to serialize with indent make sure to follow these guidelines:
//   - First argument -> prefix
//   - Second argument -> indent
func (ic ItemComputeCI) Serialize(params ...string) (serialization.Serialized, error) {
	if len(params) > 1 {
		prefix := params[0]
		indent := params[1]

		return json.MarshalIndent(ic, prefix, indent)
	}

	return json.Marshal(ic)
}
