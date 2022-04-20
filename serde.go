package jsonserde

import (
	"encoding/json"
)

// Convert ...
func Convert(data []byte) (string, error) {
	var obj []map[string]interface{}
	var s string

	if err := json.Unmarshal(data, &obj); err != nil {
		return "", err
	}

	for _, v := range obj {
		b, _ := json.Marshal(v)
		s = s + string(b)
	}

	return s, nil
}
