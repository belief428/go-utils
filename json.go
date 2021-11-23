package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
)

// AnyToJSON ToJson
func AnyToJSON(any interface{}) string {
	bytes, _ := json.Marshal(any)
	return string(bytes)
}

// AnyToByte ToByte
func AnyToByte(any interface{}) []byte {
	bytes, _ := json.Marshal(any)
	return bytes
}

// FromJSON get value from JSON string
func FromJSON(data string, v interface{}) error {
	if data == "" {
		return errors.New("data nil")
	}

	if err := json.Unmarshal([]byte(data), v); err != nil {
		return err
	}
	return nil
}

// FromJSONToMap get mapvalue from JSON string
func FromJSONToMap(data string) map[string]interface{} {
	var jsonBody map[string]interface{}
	//解析 body
	if len(data) > 0 {
		FromJSON(data, &jsonBody)
	} else {
		jsonBody = make(map[string]interface{}, 0)
	}
	return jsonBody
}

// FromJSONBytes get value from JSON bytes
func FromJSONBytes(data []byte, v interface{}) error {
	if len(data) <= 0 {
		return errors.New("data nil")
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

// FromJSONReader get value from JSON Reader
func FromJSONReader(data io.Reader, v interface{}) error {
	if data == nil {
		return errors.New("jsonFile nail")
	}

	jsonDecoder := json.NewDecoder(data)
	return jsonDecoder.Decode(v)
}

// FromJSONFile get value from JSON file
func FromJSONFile(jsonFile string, v interface{}) error {
	if len(jsonFile) <= 0 {
		return errors.New("jsonFile nail")
	}

	file, err := os.Open(jsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonDecoder := json.NewDecoder(bufio.NewReader(file))
	return jsonDecoder.Decode(v)
}
