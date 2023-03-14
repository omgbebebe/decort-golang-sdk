package serialization

import (
	"os"
)

type Writable interface {
	WriteToFile(string) error
}

type Serialized []byte

// WriteToFile writes serialized data to specified file.
//
// Make sure to use .json extension for best compatibility.
func (s Serialized) WriteToFile(path string) error {
	return os.WriteFile(path, s, 0600)
}
