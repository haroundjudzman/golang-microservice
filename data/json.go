package data

import (
	"encoding/json"
	"io"
)

// ToJSON serialises given interface into JSON format.
// The result is written into io.writer.
func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)

	return encoder.Encode(i)
}
