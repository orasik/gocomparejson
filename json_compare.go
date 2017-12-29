package jsoncompare

import (
	"encoding/json"
	"reflect"
)

// CompareJSON will check if two json strings are equal
func CompareJSON(json1, json2 string) (bool, error) {

	var jsonVal1 interface{}
	var jsonVal2 interface{}

	err1 := json.Unmarshal([]byte(json1), &jsonVal1)

	if err1 != nil {
		return false, err1
	}

	err2 := json.Unmarshal([]byte(json2), &jsonVal2)

	if err2 != nil {
		return false, err2
	}

	if reflect.DeepEqual(jsonVal1, jsonVal2) {
		return true, nil
	}

	return false, nil
}
