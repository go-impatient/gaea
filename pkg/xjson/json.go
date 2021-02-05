package xjson

import (
	jsoniter "github.com/json-iterator/go"
)

// Unmarshal .
func Unmarshal(input []byte, data interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(input, &data)
}

// Marshal .
func Marshal(data interface{}) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(&data)
}
