package grid

import (
	"encoding/json"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/serialization"
)

// Serialize returns JSON-serialized []byte. Used as a wrapper over json.Marshal and json.MarshalIndent functions.
//
// In order to serialize with indent make sure to follow these guidelines:
//   - First argument -> prefix
//   - Second argument -> indent
func (lg ListGrids) Serialize(params ...string) (serialization.Serialized, error) {
	if len(lg.Data) == 0 {
		return []byte{}, nil
	}

	if len(params) > 1 {
		prefix := params[0]
		indent := params[1]

		return json.MarshalIndent(lg, prefix, indent)
	}

	return json.Marshal(lg)
}

// Serialize returns JSON-serialized []byte. Used as a wrapper over json.Marshal and json.MarshalIndent functions.
//
// In order to serialize with indent make sure to follow these guidelines:
//   - First argument -> prefix
//   - Second argument -> indent
func (rg RecordGrid) Serialize(params ...string) (serialization.Serialized, error) {
	if len(params) > 1 {
		prefix := params[0]
		indent := params[1]

		return json.MarshalIndent(rg, prefix, indent)
	}

	return json.Marshal(rg)
}
