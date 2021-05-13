package data

import (
	"encoding/json"
	"io"
)

// ToJSON serialises given interface into JSON format.
// The result is written into io.Writer.
func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

// FromJSON deserialises JSON body from io.Reader
// into given interface.
func FromJSON(i interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(i)
}
