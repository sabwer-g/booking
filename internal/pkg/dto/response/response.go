package response

import (
	"encoding/json"

	"booking/internal/pkg/errors"
)

type Response struct {
	Success int               `json:"success"`
	Meta    json.RawMessage   `json:"meta"`
	Data    json.RawMessage   `json:"data"`
	Error   *errors.HTTPError `json:"error,omitempty"`
}
