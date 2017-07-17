// go2hal v0.2.0
// Copyright (c) 2016 Patrick Moule
// License: MIT

package hal

import (
	"encoding/json"
)

// Encoder to encode a Resource Object into a valid HAL document.
type Encoder interface {
	ToJSON(resource Resource) ([]byte, error)
}

type standardEncoder struct {
}

// NewEncoder creates a JSON encoder
func NewEncoder() Encoder {
	return new(standardEncoder)
}

// ToJSON generates a HAL document from given Resource Object.
func (enc *standardEncoder) ToJSON(resource Resource) ([]byte, error) {
	resourceObject := resource.(*resourceObject)
	namedMap := resourceObject.ToMap()

	return json.Marshal(namedMap.Content)
}
