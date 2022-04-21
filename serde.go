package jsonserde

import (
	"bytes"
	"encoding/json"
)

// Convert is a function that converts json array to string which only
// contains json objects.
// Example:
// 		json: [{"a":1},{"a":2}] -> the data that you want to convert
// 		return: {"a":1}{"a":2} -> the data that will return
func Convert(data []byte) (string, error) {
	var obj []map[string]interface{}
	var str bytes.Buffer

	if err := json.Unmarshal(data, &obj); err != nil {
		return "", err
	}

	for _, v := range obj {
		b, err := json.Marshal(v)

		if err != nil {
			continue
		}

		str.Write(b)
	}

	return str.String(), nil
}
