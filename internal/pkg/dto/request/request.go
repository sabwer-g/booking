package request

import (
	"encoding/json"
)

type Request struct {
	Meta json.RawMessage `json:"meta"`
	Data json.RawMessage `json:"data"`
}
