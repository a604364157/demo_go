package utils

import (
	"encoding/json"
	"io"
)

func JsonParse(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
